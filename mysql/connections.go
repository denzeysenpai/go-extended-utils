package mysql

import (
	"database/sql"
	"fmt"
)

func OpenConnection(pool Mysql_pool) (*sql.DB, error) {
	dsn := pool.db_user + ":" + pool.db_password + "@tcp(" + pool.db_host + ":" + pool.db_port + ")/" + pool.db_name
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Print("Connection failed! Terminating process..." + "\n")
		return conn, err
	}
	return conn, nil
}
func Execute(mysql_connections *Mysql_pool, query string, query_identifier string) ([]map[string]any, error) {
	mysql_connection, _ := OpenConnection(*mysql_connections)
	defer mysql_connection.Close()

	tx, err := mysql_connection.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mysql_connection.SetMaxOpenConns(25)
	mysql_connection.SetMaxIdleConns(25)
	mysql_connection.SetConnMaxLifetime(0)

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]any
	columnValues := make([]any, len(columns))
	columnPointers := make([]any, len(columns))

	for i := range columnValues {
		columnPointers[i] = &columnValues[i]
	}

	for rows.Next() {
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		rowObject := make(map[string]any)
		for i, colName := range columns {
			val := columnValues[i]

			if b, ok := val.([]byte); ok {
				rowObject[colName] = string(b)
			} else {
				rowObject[colName] = val
			}
		}

		results = append(results, rowObject)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return results, nil
}

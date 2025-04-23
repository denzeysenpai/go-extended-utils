package mysql

import (
	"os"
)

type Mysql_pool struct {
	db_host     string
	db_name     string
	db_user     string
	db_password string
	db_port     string
}

type Database struct {
	pool Mysql_pool
}

func CreateNewPool(_db_name string) Mysql_pool {
	return Mysql_pool{
		db_host:     os.Getenv("DB_HOST"),
		db_name:     _db_name,
		db_user:     os.Getenv("DB_USERNAME"),
		db_password: os.Getenv("DB_PASSWORD"),
		db_port:     os.Getenv("DB_PORT"),
	}
}

func CreateNewPools(db_names ...string) map[string]*Mysql_pool {
	var connections map[string]*Mysql_pool = make(map[string]*Mysql_pool)
	for _, db_name := range db_names {
		pool := CreateNewPool(db_name)
		connections[db_name] = &pool
	}
	return connections
}

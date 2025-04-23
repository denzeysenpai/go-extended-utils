package mysql

import (
	"fmt"
	"strings"
)

func GenerateInsertQuery(tableName string, data map[string]any) string {
	columns := []string{}
	values := []string{}

	// Iterate over the map
	for column, value := range data {
		columns = append(columns, fmt.Sprintf("%v", column))
		switch v := value.(type) {
		case string:
			if (strings.Contains(v, "(") && strings.Contains(v, ")")) || v == "NULL" { // Don't insert as string if it is a function
				values = append(values, fmt.Sprintf("%v", v))
			} else {
				values = append(values, fmt.Sprintf("'%v'", v))
			}
		default:
			values = append(values, fmt.Sprintf("%v", v))
		}
	}

	columnsPart := strings.Join(columns, ", ")
	valuesPart := strings.Join(values, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, columnsPart, valuesPart)
	return query
}

// returns the string NULL if str is nil
func CheckForNull(str any) any {
	if str == nil {
		return "NULL"
	} else {
		return str
	}
}

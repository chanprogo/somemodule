package mysqlsql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// OpenSQL ...
func OpenSQL(user, password, host, port, database string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database))
	if err != nil {
		panic(err)
	}
	return db
}

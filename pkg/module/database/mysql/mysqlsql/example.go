package mysqlsql

import "database/sql"

func GetNames(db *sql.DB) []string {
	l := []string{}
	var name *string

	var err error
	var rows *sql.Rows

	sqlStr := "SELECT name from example_table " + "where id>0"
	rows, err = db.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&name)
			if err == nil && name != nil {
				l = append(l, *name)
				name = nil
			}
		}
	}

	if rows != nil {
		rows.Close()
	}
	return l
}

/*
	var numInSQL uint64
	row := db.QueryRow("SELECT MAX(theNumber) FROM transaction;")
	err = row.Scan(&numInSQL)
	if err != nil {
		panic(err.Error())
	}
*/

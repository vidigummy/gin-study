package database

import "database/sql"

func database() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/TEST")
	if err != nil {
		panic(err)
	}
	return db
}

// func DoQuery(query string) (string, error) {

// }

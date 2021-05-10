package server

import (
	"database/sql"
	_ "mysql-master"
)

func Connection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "admin"
	dbName := "pizza_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

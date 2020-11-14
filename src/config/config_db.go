package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDb() (*sql.DB, error) {

	sqlDriver := "mysql"
	dbName := "go_store"
	dbUser := "root"
	dbPass := ""

	db, err := sql.Open(sqlDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")

	if err != nil {
		return nil, err
	}
	return db, nil
}

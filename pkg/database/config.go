package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func GetConn() (*sql.DB, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to initialize .env -> ", err)
	}

	sqlDriver := os.Getenv("SQL_DRIVER")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")

	db, err := sql.Open(sqlDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName))

	if err != nil {
		return nil, err
	}

	return db, nil
}

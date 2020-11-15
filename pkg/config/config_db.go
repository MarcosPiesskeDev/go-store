package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func GetDb() (*sql.DB, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to initialize .env -> ", err)
	}

	sqlDriver := os.Getenv("SQL_DRIVER")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")

	db, err := sql.Open(sqlDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")

	if err != nil {
		return nil, err
	}
	return db, nil
}

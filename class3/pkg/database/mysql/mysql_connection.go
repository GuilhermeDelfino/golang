package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabaseConnection() *sql.DB {
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")

	jdbcUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	log.Print(jdbcUrl)
	conn, err := sql.Open("mysql", jdbcUrl)

	if err != nil {
		panic(err.Error())
	}

	return conn
}

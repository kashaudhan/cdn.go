package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB


func InitDB() {
	var (
		host     = os.Getenv("PG_HOST")
		port, _  = strconv.Atoi(os.Getenv("PG_PORT"))
		user     = os.Getenv("PG_USERNAME")
		password = os.Getenv("PG_PASSWORD")
		dbname   = os.Getenv("PG_DATABASE")
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err := sql.Open("postgres", connStr)

	if err != nil {
		panic("DB connection failed")
	}

	DB.SetConnMaxIdleTime(5)
	CreateTables(DB)
}

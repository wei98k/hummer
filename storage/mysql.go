package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"hummer/config"
	"log"
)

var DB *sql.DB

func InitMySQL() {
	var err error
	dsn := config.DBConfig.DSN()
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}
}

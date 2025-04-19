package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitMySQL() {
	var err error
	dsn := "root:boy.root1231M@tcp(127.0.0.1:33061)/hummer?parseTime=true"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}
}

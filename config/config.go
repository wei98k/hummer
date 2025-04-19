package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var DBConfig MySQLConfig
var APIToken string

func (cfg MySQLConfig) DSN() string {
	return cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.DBName + "?parseTime=true"
}

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
	log.Println("打印配置信息")
	DBConfig = MySQLConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DB"),
	}
	APIToken = os.Getenv("API_TOKEN")
}

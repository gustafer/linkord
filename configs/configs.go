package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LoadPort() string {
	return os.Getenv("PORT")
}

func LoadDbConnStr() string {
	var (
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbname   = os.Getenv("DB_NAME")
	)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	return connStr
}

func GetRedirectUrl() string {
	return os.Getenv("REDIRECT_URL")
}

func LoadAuthKey() string {
	return os.Getenv("AUTH_KEY")
}

package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB


func init() {

	var err error

    err = godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    DB, err = sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/" + dbName)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

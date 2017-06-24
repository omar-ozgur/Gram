package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/omar-ozgur/gram/utilities"
	"os"
	_ "time"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "gram"
	DB_HOST     = "localhost"
)

var DB *sql.DB

func InitDB() {
	DBInfo := os.Getenv("GRAM_DB_INFO")
	if DBInfo == "" {
		DBInfo = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
			DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	}
	var err error
	DB, err = sql.Open("postgres", DBInfo)
	utilities.CheckErr(err)

	_, err = DB.Exec("SELECT * FROM users")
	if err != nil {
		_, err = DB.Exec(`CREATE TABLE users (
           id SERIAL,
           first_name text,
           last_name text,
           email text,
           fb_id text,
           password bytea,
           time_created timestamp DEFAULT now()
           );`)
		utilities.CheckErr(err)
	}
}

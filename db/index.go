package db

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/omar-ozgur/gram/utilities"
	"io/ioutil"
	"os"
	_ "time"
)

type DBParams struct {
	User     string
	Password string
	Name     string
	Host     string
	SSLMode  string
	Info     string
}

const (
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_NAME     = "gram"
	DB_HOST     = "localhost"
)

var DB *sql.DB

func FindDBParam(param *string, alias string, env string, defaultValue string) {
	scanner := bufio.NewScanner(os.Stdin)

	*param = os.Getenv(env)
	if *param != "" {
		fmt.Printf("The currently saved DB %s '%s'. Type a new %s, or press enter to keep it.\n", alias, *param, alias)
	} else {
		*param = defaultValue
		fmt.Printf("The default DB %s is '%s'. Type a new %s, or press enter to keep it.\n", alias, *param, alias)
	}

	scanner.Scan()
	input := scanner.Text()
	if input != "" {
		*param = input
	}

	os.Setenv(env, *param)
}

func FindDBInfo(dbParams *DBParams) {
	dbParams.Info = os.Getenv("GRAM_DB_INFO")
	if dbParams.Info != "" {
		return
	}

	fmt.Println("No database information was found. Please provide your database credentials. If the database does not exist, you will be prompted to create it.")

	FindDBParam(&dbParams.User, "username", "GRAM_DB_USER", utilities.DefaultDBUser)
	FindDBParam(&dbParams.Password, "password", "GRAM_DB_PASSWORD", utilities.DefaultDBPassword)
	FindDBParam(&dbParams.Name, "name", "GRAM_DB_NAME", utilities.DefaultDBName)
	FindDBParam(&dbParams.Host, "host", "GRAM_DB_HOST", utilities.DefaultDBHost)
	FindDBParam(&dbParams.SSLMode, "SSL mode", "GRAM_DB_SSLMODE", utilities.DefaultDBSSLMode)

	dbParams.Info = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		dbParams.User, dbParams.Password, dbParams.Name, dbParams.Host)
	os.Setenv("GRAM_DB_INFO", dbParams.Info)
}

func InitDB(service string) {
	var dbParams DBParams

	data, err := ioutil.ReadFile("config/dbParams.json")
	if err == nil {
		json.Unmarshal(data, &dbParams)
	}

	if dbParams.Info == "" {
		FindDBInfo(&dbParams)
	}

	json, err := json.Marshal(dbParams)
	utilities.CheckErr(err)

	err = ioutil.WriteFile("config/dbParams.json", json, 0644)
	utilities.CheckErr(err)

	DB, err := sql.Open("postgres", dbParams.Info)
	if err != nil {
		panic(fmt.Sprintf("Error: An error occurred while opening the SQL database\n%v", err))
	}

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
	utilities.CheckErr(err)
}

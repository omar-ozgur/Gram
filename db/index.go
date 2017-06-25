package db

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/omar-ozgur/gram/utilities"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"time"
)

type DBParams struct {
	User     string
	Password string
	Name     string
	Host     string
	SSLMode  string
}

var DB *sql.DB
var DBInfo string
var GCM cipher.AEAD
var Nonce []byte

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func FindDBParam(param *string, alias string, defaultValue string, encrypt bool) {
	scanner := bufio.NewScanner(os.Stdin)

	if !encrypt {
		if *param != "" {
			fmt.Printf("The currently saved DB %s is '%s'. Type a new %s, or press enter to keep it.\n", alias, *param, alias)
		} else {
			*param = defaultValue
			fmt.Printf("The default DB %s is '%s'. Type a new %s, or press enter to keep it.\n", alias, *param, alias)
		}
	} else {
		if *param != "" {
			fmt.Printf("Please enter a new %s, or press enter to keep your old one.\n", alias)
		} else {
			fmt.Printf("Please enter a new %s.\n", alias)
		}
	}

	complete := false
	for complete == false {
		scanner.Scan()
		input := scanner.Text()
		if input != "" {
			if encrypt {
				encryptedInput := GCM.Seal(nil, Nonce, []byte(input), nil)
				*param = hex.EncodeToString(encryptedInput)
			} else {
				*param = input
			}
			complete = true
		} else if encrypt && *param == "" {
			fmt.Printf("The %s cannot be blank. Please try again.\n", alias)
		} else {
			complete = true
		}
	}
}

func FindDBInfo(dbParams *DBParams) {
	DBInfo = os.Getenv("GRAM_DB_INFO")
	if DBInfo != "" {
		return
	}

	key := os.Getenv("GRAM_ENCRYPTION_KEY")
	if key == "" {
		fmt.Println("Error: No database encryption key was found.")
		newKey := RandStringBytesMaskImprSrc(16)
		fmt.Println("Please set the GRAM_ENCRYPTION_KEY environment variable to the following 16-byte value, or generate your own.")
		fmt.Println(string(newKey))
		os.Exit(1)
	}

	c, err := aes.NewCipher([]byte(key))
	utilities.CheckErr(err)

	GCM, err = cipher.NewGCM(c)
	utilities.CheckErr(err)

	Nonce = make([]byte, GCM.NonceSize())

	count := 0
	v := reflect.ValueOf(*dbParams)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != "" {
			count++
		}
	}

	if count == 0 {
		fmt.Println("No database information was found. Please provide your database credentials. If the database does not exist, you will be prompted to create it.")
	}

	if count < v.NumField() {
		FindDBParam(&dbParams.User, "username", utilities.DefaultDBUser, false)
		FindDBParam(&dbParams.Password, "password", "", true)
		FindDBParam(&dbParams.Name, "name", utilities.DefaultDBName, false)
		FindDBParam(&dbParams.Host, "host", utilities.DefaultDBHost, false)
		FindDBParam(&dbParams.SSLMode, "SSL mode", utilities.DefaultDBSSLMode, false)
	}

	decodedHex, err := hex.DecodeString(dbParams.Password)
	utilities.CheckErr(err)

	plainPassword, err := GCM.Open(nil, Nonce, decodedHex, nil)
	utilities.CheckErr(err)

	DBInfo = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		dbParams.User, string(plainPassword), dbParams.Name, dbParams.Host)
}

func InitDB(service string) {
	var dbParams DBParams

	data, err := ioutil.ReadFile("config/dbParams.json")
	if err == nil {
		json.Unmarshal(data, &dbParams)
	}

	FindDBInfo(&dbParams)

	json, err := json.Marshal(dbParams)
	utilities.CheckErr(err)

	err = ioutil.WriteFile("config/dbParams.json", json, 0644)
	utilities.CheckErr(err)

	DB, err := sql.Open("postgres", DBInfo)
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

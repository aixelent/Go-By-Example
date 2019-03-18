package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

const (
	DbDriver = "mysql"
	DbName   = "testing"
	DbUser   = "user"
	DbPass   = "userpass"
	DbURL    = DbUser + ":" + DbPass + "@/" + DbName
)

func logErr(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
}

func init() {
	db, err = sql.Open(DbDriver, DbURL)
	logErr(err)

	if err = db.Ping(); err != nil {
		logErr(err)
	} else {
		fmt.Println("DB Connected successfully.")
	}
}

func main() {
	defer db.Close()
}

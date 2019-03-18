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
	DbName   = "test"
	DbUser   = "test"
	DbPass   = "test"
	DbURL    = DbUser + ":" + DbPass + "@/" + DbName
)

func init() {
	db, err = sql.Open(DbDriver, DbURL)
	logErr(err)

	err = db.Ping()
	logErr(err)
}

func logErr(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}

func InsertQuery() {
	query, err := db.Prepare("INSERT artist SET artist=?, alias=?, created=?")
	logErr(err)

	exec, err := query.Exec("Brock", "Bvdub", "2019-02-20")
	logErr(err)

	id, err := exec.LastInsertId()
	logErr(err)

	fmt.Println(id)
}

func SelectData() {
	rows, err := db.Query("SELECT * FROM artist")
	logErr(err)

	for rows.Next() {
		var uid int
		var artist string
		var alias string
		var created string
		err := rows.Scan(&uid, &artist, &alias, &created)
		logErr(err)

		fmt.Println("ID:", uid, "\nUser:", artist, "\nDepartname:", alias, "\nCreated:", created)
	}
}

func main() {
	defer db.Close()
	InsertQuery()
	SelectData()
}

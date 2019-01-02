package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func logPanic(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
		panic(err)
	}
}

func connection() *sql.DB {
	connection := "postgres://postgres:macro1@localhost:5432/testdb?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	logPanic(err)

	err = db.Ping()
	logPanic(err)

	return db
}

func main() {
	db := connection()
	defer db.Close()

	result, err := db.Query("SELECT * FROM album a")
	logPanic(err)
	defer result.Close()

	columns, err := result.Columns()
	logPanic(err)
	fmt.Printf("Selected columns: %v", columns)

	colType, err := result.ColumnTypes()
	logPanic(err)
	for _, col := range colType {
		fmt.Printf("%+v\n", col)
	}
}

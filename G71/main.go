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

func main() {
	connection := "postgres://postgres:macro1@localhost:5432/testdb?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	logPanic(err)

	defer db.Close()
	err = db.Ping()
	logPanic(err)

	fmt.Println("Pinged, ok!")
}

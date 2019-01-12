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

type Album struct {
	Author sql.NullString
	Title  sql.NullString
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

	result, err := db.Query("SELECT author, title FROM album;")
	logPanic(err)
	defer result.Close()

	albums := []Album{}
	for result.Next() {
		if result.Err() != nil {
			panic(result.Err())
		}
		a := Album{}
		if err := result.Scan(&a.Author, &a.Title); err != nil {
			panic(err)
		}
		albums = append(albums, a)
	}
	fmt.Printf("Retrieved albums: %+v\n", albums)
}

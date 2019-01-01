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

	_, err := db.Exec("TRUNCATE TABLE album;")
	logPanic(err)
	fmt.Println("Table truncated")

	result, err := db.Exec("INSERT INTO album(ID, TITLE, AUTHOR, GENRE, STYLE) VALUES (1, 'At Night This City Becomes The Sea', 'Bvdub', 'Electronic', 'Ambient'), (2, 'Strength In Solitude LP', 'Bvdub', 'Electronic', 'Ambient'), (3, 'A Prayer To False Gods', 'Bvdub', 'Electronic', 'Ambient'), (4, 'We Were The Sun', 'Bvdub', 'Electronic', 'Ambient'), (5, 'A Silent Reign', 'Bvdub', 'Electronic', 'Ambient'), (6, 'Presents Deep Space Mix 21', 'Bvdub', 'Electronic', 'Ambient'), (7, 'Resistance Is Beautiful', 'Bvdub', 'Electronic', 'Ambient'), (8, 'Serenity', 'Bvdub', 'Electronic', 'Ambient'), (9, 'One Last Look At The Sea', 'Bvdub', 'Electronic', 'Ambient'), (10, 'Born In Tokyo', 'Bvdub', 'Electronic', 'Ambient'), (11, 'At Night This City Becomes The Sea', 'Bvdub', 'Electronic', 'Ambient'), (12, 'The Art Of Dying Alone', 'Bvdub', 'Electronic', 'Ambient') ")

	logPanic(err)

	rowsAffected, err := result.RowsAffected()
	logPanic(err)
	fmt.Printf("Inserted correctly: %d rows", rowsAffected)

	rows, err := db.Query("SELECT * FROM album;")
	logPanic(err)

	count := 0
	for rows.Next() {
		count++
	}
	fmt.Printf("\n%d rows was selected", count)
}

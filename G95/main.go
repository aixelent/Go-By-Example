package main

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var tpl *template.Template

type User struct {
	ID         int64
	Password   []byte
	ArtistName string
	RealName   string
	Aliases    string
	Profile    string
	Variations string
}

func init() {
	db, err = sql.Open("mysql", "test:test@/test")
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func logErr(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {
	defer db.Close()
	http.HandleFunc("/", index)
	http.HandleFunc("/userForm", userForm)
	http.HandleFunc("/createUsers", createUsers)
	http.HandleFunc("/editUsers", editUsers)
	http.HandleFunc("/deleteUsers", deleteUsers)
	http.HandleFunc("/updateUsers", updateUsers)
	log.Println("Server is up on 8080.")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func userForm(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "userForm.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := User{}
		user.Username = r.FormValue("userName")
		user.FirstName = r.FormValue("firstName")
		user.Email = r.FormValue("lastName")
		bcryptPass, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.Password = bcryptPass
		_, err = db.Exec(`INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)`,
			user.Username, user.FirstName, user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("UPDATE users SET username = ?, first_name = ?, last_name = ?, WHERE id = ?",
		r.FormValue("userName"),
		r.FormValue("firstName"),
		r.FormValue("lastName"),
		r.FormValue("id"),
	)
	logErr("Update:", err)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Bad ID", http.StatusBadRequest)
	}
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	logErr("Delete:", err)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editUsers(w http.ResponseWriter, r *http.Request) {

}

func index(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM users;`)
	if err != nil {
		log.Println("Tutej:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		rows.Scan(&u)
		users = append(users, u)
	}
	log.Println(users)
	tpl.ExecuteTemplate(w, "index.gohtml", users)
}

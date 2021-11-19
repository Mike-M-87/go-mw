package main

import "time"
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func hellopage() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
}

func FileServer() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static", http.StripPrefix("/static", fs))
}

func muxrouter() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
}

func simple_sql() {

	db, err := sql.Open("mysql", "golden:password@(127.0.0.1:3306)/staff?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	//create table
	{
		query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
	//insert a user
	{
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		fmt.Println(id)
	}
	//Query a row
	{
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username, password, createdAt)
	}
	//select all rows
	{
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id,username,password,created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close() // I'm not sure what happens here lol
		var users []user

		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", users)

	}
	//Deleting a row
	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
		if err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	http.ListenAndServe(":8081", nil)

}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	PORT        = os.Getenv("PORT")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_DATABASE = os.Getenv("DB_DATABASE")
	DB          *sql.DB
)

func main() {
	DB, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE))
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	if PORT == "" {
		log.Printf("port not specified!")
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/books", func(w http.ResponseWriter, req *http.Request) {
		rows, err := DB.Query(`select book.id, book.title, genre.genre_name, author.author_name from book
		join author on author.id = book.author_id
		join book_genre on book_genre.book_id = book.id
		join genre on genre.id = book_genre.genre_id;`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		resp := []map[string]interface{}{}
		for rows.Next() {
			var id int
			var title string
			var genre string
			var author string
			err := rows.Scan(&id, &title, &genre, &author)
			if err != nil {
				log.Fatal(err)
			}
			row := map[string]interface{}{
				"id":     id,
				"title":  title,
				"genre":  genre,
				"author": author,
			}
			resp = append(resp, row)
		}
		body, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}).Methods("GET")

	router.HandleFunc("/genres", func(w http.ResponseWriter, req *http.Request) {
		rows, err := DB.Query(`select id, genre_name from genre order by id`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		resp := []map[string]interface{}{}
		for rows.Next() {
			var id int
			var genre string
			err := rows.Scan(&id, &genre)
			if err != nil {
				log.Fatal(err)
			}
			row := map[string]interface{}{
				"id":    id,
				"genre": genre,
			}
			resp = append(resp, row)
		}
		body, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}).Methods("GET")

	router.HandleFunc("/authors", func(w http.ResponseWriter, req *http.Request) {
		rows, err := DB.Query(`select id, author_name from author order by id`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		resp := []map[string]interface{}{}
		for rows.Next() {
			var id int
			var author string
			err := rows.Scan(&id, &author)
			if err != nil {
				log.Fatal(err)
			}
			row := map[string]interface{}{
				"id":     id,
				"author": author,
			}
			resp = append(resp, row)
		}
		body, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}).Methods("GET")

	router.HandleFunc("/authors", func(w http.ResponseWriter, req *http.Request) {
		type Author struct {
			Name string `json:"name"`
		}
		var author Author

		err := json.NewDecoder(req.Body).Decode(&author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		DB.Exec(fmt.Sprintf("INSERT INTO author (author_name) VALUES ('%s')", author.Name))
	}).Methods("POST")

	router.HandleFunc("/genres", func(w http.ResponseWriter, req *http.Request) {
		type Genre struct {
			Name string `json:"name"`
		}
		var genre Genre

		err := json.NewDecoder(req.Body).Decode(&genre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		DB.Exec(fmt.Sprintf("INSERT INTO genre (genre_name) VALUES ('%s')", genre.Name))
	}).Methods("POST")

	router.HandleFunc("/books", func(w http.ResponseWriter, req *http.Request) {
		type Book struct {
			Title    string `json:"title"`
			GenreID  int    `json:"genre_id"`
			AuthorID int    `json:"author_id"`
		}
		var book Book

		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		DB.Exec(fmt.Sprintf("INSERT INTO book (title, author_id) VALUES ('%s', %d)", book.Title, book.AuthorID))
		DB.Exec(fmt.Sprintf(`INSERT INTO book_genre (book_id, genre_id) VALUES (
			(SELECT id from book where title = '%s'),
			%d
		);`, book.Title, book.GenreID))
	}).Methods("POST")

	log.Printf("running on port: %s\n", PORT)
	http.ListenAndServe(":"+PORT, router)
}

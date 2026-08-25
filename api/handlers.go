package api

import (
	"encoding/json"
	"fmt"
	"goapi/db"
	"net/http"
)

// GetUsers retourne la liste des utilisateurs
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, username, email FROM user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []db.User
	for rows.Next() {
		var u db.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT b.id, b.isbn13, b.title, b.subtitle, b.kind, b.publisheddate,p.name FROM book as b	INNER JOIN publisher p on p.id=b.publisherid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []db.Book
	for rows.Next() {
		var b db.Book
		if err := rows.Scan(&b.ID, &b.ISBN13, &b.Title, &b.Subtitle, &b.Kind, &b.PublishedDate, &b.Publisher); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, b)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	fmt.Println(books)
}

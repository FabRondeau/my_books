package handlers

import (
	"encoding/json"
	"fmt"
	"goapi/db"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Fix empty password when Signin up (because of password string json: "-")
	var req db.Book
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"success": false, "message": "Données invalides"}`, http.StatusBadRequest)
		return
	}
	// Crée un book avec les données de la requête
	book := db.Book{
		ISBN13:        req.ISBN13,
		Title:         req.Title,
		Kind:          req.Kind,
		Subtitle:      req.Subtitle,
		Publisher:     req.Publisher,
		PublishedDate: req.PublishedDate,
	}

	// Recherche l'éditeur
	rows, err := db.DB.Query(
		"SELECT id,name FROM publisher where name = ?",
		book.Publisher,
	)
	if err != nil {
		fmt.Println("Erreur")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	fmt.Println("SQL ok")
	if rows.Next() {
		var p db.Publisher
		fmt.Println(p)
		// for rows.Next() {
		// var p db.Publisher
		fmt.Println("rows next")
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rows.Close()
		// }

	} else {
		// TODO: créer l'éditeur puisqu'il n'existe pas.
		InsertPublisher(w, book.Publisher)
	}
	// TODO: CHECK

}

func InsertBook(w http.ResponseWriter, p db.Publisher, b db.Book) {
	fmt.Println("Insert the book")
	_, err := db.DB.Exec(
		"INSERT INTO book (isbn13, title, subtitle, kind, publisherid,publisheddate) VALUES (?, ?, ?, ?, ?, ?)",
		b.ISBN13, b.Title, b.Subtitle, b.Kind, p.ID, b.PublishedDate,
	)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Erreur lors de la création du livre"}`, http.StatusConflict)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"book":    b,
	})

}

func InsertPublisher(w http.ResponseWriter, publishername string) {
	fmt.Println("Insert publisher", publishername)
	_, err := db.DB.Exec(
		"INSERT INTO publisher (name) VALUES (?)",
		publishername,
	)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Erreur lors de la création du livre"}`, http.StatusConflict)
		return
	}
	// Return the book
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   true,
		"publisher": publishername,
	})

}

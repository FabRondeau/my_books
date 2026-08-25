package db

import (
	"database/sql"
)

// RunMigrations exécute les migrations nécessaires pour la base de données
func RunMigrations(db *sql.DB) error {
	// Crée la table `users`
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS user (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					username VARCHAR(150) UNIQUE NOT NULL,
					email VARCHAR(150) UNIQUE NOT NULL,
					password VARCHAR(60) NOT NULL
        );
				CREATE TABLE IF NOT EXISTS book (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					isbn13 TEXT UNIQUE NOT NULL,
					kind TEXT NOT NULL,
					totalitems INT NOT NULL,
					title VARCHAR(255) NOT NULL,
					subtitle VARCHAR(255) NOT NULL,
					publisheddate VARCHAR(15) NOT NULL,
					publisherid INTEGER NOT NULL,
					FOREIGN KEY (publisherid) REFERENCES publisher(id)
        );
				CREATE TABLE IF NOT EXISTS author (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					fullname VARCHAR(150) NOT NULL
        );
				CREATE TABLE IF NOT EXISTS publisher (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					name VARCHAR(150) NOT NULL
        );
				CREATE TABLE IF NOT EXISTS book_author(
					bookid INTEGER NOT NULL,
					authorid INTEGER NOT NULL,
					FOREIGN KEY (bookid) REFERENCES book(id),
					FOREIGN KEY (authorid) REFERENCES author(id)
				);
				CREATE TABLE IF NOT EXISTS user_location (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					userid INTEGER NOT NULL,
					location VARCHAR NOT NULL,
					FOREIGN KEY (userid) REFERENCES user(id)
				);
				CREATE TABLE IF NOT EXISTS book_user_location (
					userlocationid INTEGER NOT NULL,
					bookid INTEGER NOT NULL,
					FOREIGN KEY (userlocationid) REFERENCES user_location(id),
					FOREIGN KEY (bookid) REFERENCES book(id)
				);


				-- INSERT INTO author (fullname) values ('Jérôme Loubry');
				-- INSERT INTO publisher (name) values ('Le livre de poche');
				-- INSERT INTO book (isbn13,kind,totalitems,title,subtitle,publisheddate,publisherid) values ('9782253181590','Roman',1,'Les Refuges','','2019',1);
				-- INSERT INTO book_author (bookid,authorid) values (1,1);


    `)
	if err != nil {
		return err
	}

	// Ajoute d'autres migrations ici si nécessaire
	return nil
}

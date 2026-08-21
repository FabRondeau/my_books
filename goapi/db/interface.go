package db

import "database/sql"

// DBInterface définit les méthodes nécessaires pour interagir avec la base de données
type DBInterface interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Ping() error
	Close() error
}

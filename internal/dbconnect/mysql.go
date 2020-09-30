package dbconnect

import "database/sql"

// DbConn for database connection
func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goassign"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:8889)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

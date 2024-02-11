package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	host  = "localhost"
	port  = 5432
	user  = "admin"
	pass  = "123"
	dbase = "rinha"
)

func getPostgresConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbase))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	return db, err
}

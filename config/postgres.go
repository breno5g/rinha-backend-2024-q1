package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	host  = os.Getenv("POSTGRES_HOST")
	port  = 5432
	user  = os.Getenv("POSTGRES_USER")
	pass  = os.Getenv("POSTGRES_PASSWORD")
	dbase = os.Getenv("POSTGRES_DB")
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

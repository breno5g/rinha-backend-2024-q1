package config

import "database/sql"

var (
	db *sql.DB
)

func Init() error {
	var err error
	db, err = getPostgresConnection()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

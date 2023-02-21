package databaseSQL

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type database struct {
	db *sql.DB
}

var instance *database

func GetInstance() (*sql.DB, error) {
	if instance == nil {
		cfg := mysql.Config{
			User:   "root",
			Passwd: "password",
			Net:    "tcp",
			Addr:   "127.0.0.1:3306",
			DBName: "my-app",
		}

		connString := cfg.FormatDSN()

		db, err := sql.Open("mysql", connString)
		if err != nil {
			return nil, fmt.Errorf("error opening database connection: %s", err)
		}
		fmt.Println("Creating single DB instance now.")
		err = db.Ping()
		if err != nil {
			return nil, fmt.Errorf("error pinging database: %s", err)
		}

		instance = &database{db: db}
	}
	fmt.Println("single DB instance already exist.")
	return instance.db, nil
}
func Close() error {
	if instance == nil {
		return nil
	}
	err := instance.db.Close()
	if err != nil {
		return fmt.Errorf("error closing database connection: %s", err)
	}
	instance = nil
	return nil
}

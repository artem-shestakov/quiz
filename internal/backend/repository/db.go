package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Address  string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func CreateDB(db *Database) (*sqlx.DB, error) {
	database, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		db.Address, db.Port, db.User, db.DBName, db.Password, db.SSLMode))
	if err != nil {
		return nil, err
	}
	if err := database.Ping(); err != nil {
		return nil, err
	}
	return database, nil
}
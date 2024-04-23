package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	ID    int
	Brand string
	Model string
	Year  string
	Price float64
}

var db *sql.DB

func InitDB(databasePath string) error {

	var err error
	db, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}

	return nil

}

func (c *Car) Create() error {

	_, err := db.Exec("INSERT INTO cars (brand, model, year, price) VALUES (?, ?, ?, ?)", c.Brand, c.Model, c.Year, c.Price)
	if err != nil {
		return err
	}
	return nil
}

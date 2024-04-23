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

// TODO
// Check empty atributes of Car

var Db *sql.DB

func InitDB(databasePath string) error {

	var err error
	Db, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}

	return nil

}

func (c *Car) Create() error {

	_, err := Db.Exec("INSERT INTO cars (brand, model, year, price) VALUES (?, ?, ?, ?)", c.Brand, c.Model, c.Year, c.Price)
	if err != nil {
		return err
	}
	return nil
}

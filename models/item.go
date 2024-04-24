package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ListCars []Car

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

func (l *ListCars) ListAll() error {

	rows, err := Db.Query("SELECT * FROM cars;")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var car Car
		err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &car.Price)
		if err != nil {
			return err
		}
		*l = append(*l, car)
	}

	return rows.Err()
}

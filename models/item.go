package models

import (
	"database/sql"
	"fmt"

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

func (c *Car) List(car_id int) (*Car, error) {

	row, err := Db.Query("SELECT * FROM cars c WHERE c.id = ?;", car_id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var carResult Car

	if !row.Next() {
		if err := row.Err(); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Car with ID %v not found", car_id)
	}
	err = row.Scan(&carResult.ID, &carResult.Brand, &carResult.Model, &carResult.Year, &carResult.Price)
	if err != nil {
		return nil, err
	}
	// for row.Next() {
	// 	err := row.Scan(&carResult.ID, &carResult.Brand, &carResult.Model, &carResult.Year, &carResult.Price)
	// 	log.Println(carResult)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return &carResult, row.Err()

}

func (c *Car) Delete(car_id int) error {

	row := Db.QueryRow("SELECT 1 FROM cars WHERE id = ?;", car_id)
	var exists bool

	err := row.Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("Car with ID %v not found", car_id)
	}

	result, err := Db.Exec("DELETE FROM cars where id = ?;", car_id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted")
	}

	return err

}

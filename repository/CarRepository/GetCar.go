package CarRepository

import (
	"context"
	"final-project-go/config"
	"final-project-go/models"
)

func GetCar(ctx context.Context, id string) (models.Car, error) {
	var car models.Car
	db, err := config.MySQL()
	defer db.Close()

	if err != nil {
		return models.Car{}, err
	}

	queryText := "SELECT id, title, price, image, description, passenger, luggage, car_type, isDriver, duration FROM cars WHERE id = ?"
	row, err := db.QueryContext(ctx, queryText, id)

	if row.Next() {
		err := row.Scan(&car.ID, &car.TITLE, &car.PRICE, &car.IMAGE, &car.DESCRIPTION, &car.PASSENGER, &car.LUGGAGE, &car.CARTYPE, &car.ISDRIVER, &car.DURATION)

		if err != nil {
			return models.Car{}, err
		}
	}

	return car, nil

}

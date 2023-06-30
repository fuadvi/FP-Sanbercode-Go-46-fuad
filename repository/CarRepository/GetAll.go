package CarRepository

import (
	"context"
	"final-project-go/config"
	"final-project-go/models"
)

func GetAll(ctx context.Context) ([]models.Car, error) {
	var cars []models.Car

	db, err := config.MySQL()

	if err != nil {
		return nil, err
	}

	queryTex := "SELECT id, title, price, image, description, passenger, luggage, car_type, isDriver, duration FROM cars"

	rows, err := db.QueryContext(ctx, queryTex)

	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.ID, &car.TITLE, &car.PRICE, &car.IMAGE, &car.DESCRIPTION, &car.PASSENGER, &car.LUGGAGE, &car.CARTYPE, &car.ISDRIVER, &car.DURATION)

		if err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	return cars, nil
}

package CarRepository

import (
	"context"
	"errors"
	"final-project-go/Request"
	"final-project-go/config"
)

func Insert(ctx context.Context, car Request.Car) error {
	db, err := config.MySQL()
	defer db.Close()

	if err != nil {
		return errors.New("can't connect to mysql")
	}

	queryText := "INSERT INTO cars (title, price, image, description, passenger, luggage, car_type, isDriver) VALUES (?,?,?,?,?,?,?,?)"
	_, err = db.ExecContext(ctx, queryText, car.TITLE, car.PRICE, car.IMAGE, car.DESCRIPTION, car.PASSENGER, car.LUGGAGE, car.CARTYPE, car.ISDRIVER)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

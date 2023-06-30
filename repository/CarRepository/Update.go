package CarRepository

import (
	"context"
	"final-project-go/Request"
	"final-project-go/config"
)

func Update(ctx context.Context, car Request.Car, id string) error {

	db, err := config.MySQL()
	defer db.Close()
	if err != nil {
		return err
	}

	queryText := "UPDATE cars SET title = ?, price = ?, image = ?, description = ?, passenger = ?, luggage = ?, car_type = ?, isDriver = ?, duration = ? WHERE id = ?"
	_, err = db.ExecContext(ctx, queryText, car.TITLE, car.PRICE, car.IMAGE, car.DESCRIPTION, car.PASSENGER, car.LUGGAGE, car.CARTYPE, car.ISDRIVER, car.DURATION, id)

	if err != nil {
		return err
	}

	return nil
}

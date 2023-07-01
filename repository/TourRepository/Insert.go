package TourRepository

import (
	"context"
	"final-project-go/Request"
	"final-project-go/config"
)

func Insert(ctx context.Context, tourReq Request.TourRequest) error {
	db, err := config.MySQL()
	defer db.Close()

	if err != nil {
		return err
	}

	queryText := "INSERT INTO tours (title, price, duration, description) VALUES (?, ?, ?, ?)"
	_, err = db.ExecContext(ctx, queryText, tourReq.TITLE, tourReq.PRICE, tourReq.DURATION, tourReq.DESCRIPTION)

	if err != nil {
		return err
	}

	return nil
}

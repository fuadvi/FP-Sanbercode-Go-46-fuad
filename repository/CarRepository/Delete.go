package CarRepository

import (
	"context"
	"final-project-go/config"
)

func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	defer db.Close()

	if err != nil {
		return err
	}

	queryText := "DELETE FROM users WHERE id = ?"
	_, err = db.ExecContext(ctx, queryText, id)

	if err != nil {
		return err
	}

	return nil
}

package UserRepostory

import (
	"context"
	"final-project-go/config"
	"final-project-go/models"
	"final-project-go/utilitis"
	"log"
)

func Insert(ctx context.Context, user models.Users) error {
	db, err := config.MySQL()
	defer db.Close()

	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}

	password, err := utilitis.HashPassword(user.PASSWORD)

	if err != nil {
		log.Fatal(err)
	}

	queryText := "INSERT INTO users (name, email, no_hp, password) VALUES (?,?,?,?)"
	_, err = db.ExecContext(ctx, queryText, user.NAME, user.EMAIL, user.NOHP, password)
	if err != nil {
		return err
	}
	return nil
}

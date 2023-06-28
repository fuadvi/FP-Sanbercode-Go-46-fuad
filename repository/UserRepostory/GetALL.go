package UserRepostory

import (
	"context"
	"final-project-go/config"
	"final-project-go/models"
	"log"
)

func GetAll(ctx context.Context) ([]models.Users, error) {
	var users []models.Users

	db, err := config.MySQL()
	defer db.Close()

	if err != nil {
		log.Fatal("Cant connect to mysql", err)
	}

	queryText := "SELECT * FROM users"
	rows, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user models.Users
		err := rows.Scan(&user.ID, &user.NAME, &user.EMAIL, &user.NOHP, &user.PASSWORD)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

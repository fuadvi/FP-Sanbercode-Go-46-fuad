package UserRepostory

import (
	"context"
	"final-project-go/config"
	"final-project-go/models"
	"final-project-go/utilitis"
	"fmt"
	"log"
)

func Login(ctx context.Context, email, password string) (models.Users, error) {
	db, err := config.MySQL()
	defer db.Close()
	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}

	var user models.Users

	queryText := "SELECT name, email, no_hp, password FROM users WHERE email = ?  LIMIT 1"
	row, err := db.QueryContext(ctx, queryText, email)

	fmt.Println(email, password)

	if row.Next() {
		err := row.Scan(&user.NAME, &user.EMAIL, &user.NOHP, &user.PASSWORD)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("success")
	} else {
		fmt.Println("gagal")
	}

	err = utilitis.ComparePasswords(user.PASSWORD, password)
	if err != nil {
		fmt.Println(user.PASSWORD, password)
		log.Fatal("Email dan Password Salah")
	}

	return user, nil
}

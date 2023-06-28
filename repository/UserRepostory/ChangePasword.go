package UserRepostory

import (
	"context"
	"final-project-go/Request"
	"final-project-go/config"
	"final-project-go/models"
	"final-project-go/utilitis"
	"fmt"
	"log"
)

func ChangePassword(ctx context.Context, request Request.ChangePasswordRequest) error {
	db, err := config.MySQL()
	defer db.Close()

	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}

	var user models.Users

	queryText := "SELECT name, email, no_hp, password FROM users WHERE email = ?  LIMIT 1"
	row, err := db.QueryContext(ctx, queryText, request.Email)

	if row.Next() {
		err := row.Scan(&user.NAME, &user.EMAIL, &user.NOHP, &user.PASSWORD)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("success")
	} else {
		fmt.Println("gagal")
	}

	err = utilitis.ComparePasswords(user.PASSWORD, request.OldPassword)
	if err != nil {
		log.Fatal("Email dan Password Salah")
	}

	password, err := utilitis.HashPassword(request.Password)
	if err != nil {
		log.Fatal(err)
	}

	queryText = "UPDATE users set password = ? where email = ?"
	_, err = db.ExecContext(ctx, queryText, password, request.Email)
	if err != nil {
		return err
	}

	return nil
}

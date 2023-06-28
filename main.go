package main

import (
	"context"
	"encoding/json"
	"final-project-go/Request"
	"final-project-go/models"
	"final-project-go/repository/UserRepostory"
	"final-project-go/utilitis"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	router := httprouter.New()
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/change-password", JWTMiddleware(ChangePassword))
	router.GET("/list-user", JWTMiddleware(ListUser))
	http.ListenAndServe(":8080", router)
}

func JWTMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r, ps)
	}
}

func CreateToken(username string) (string, error) {
	if username == "" {
		log.Fatal("Username tidak valid")
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ListUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	users, err := UserRepostory.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	utilitis.ResponseJSON(w, users, http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var user models.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utilitis.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := UserRepostory.Insert(ctx, user); err != nil {
		utilitis.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utilitis.ResponseJSON(w, res, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	var loginReq Request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Gagal membaca body permintaan", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	user, err := UserRepostory.Login(ctx, loginReq.Email, loginReq.Password)
	if err != nil {
		log.Fatal(err)
	}

	token, err := CreateToken(user.NAME)
	if err != nil {
		log.Fatal(err)
	}

	res := map[string]string{
		"status": "Successfully",
		"token":  token,
	}

	utilitis.ResponseJSON(w, res, http.StatusCreated)
}

func ChangePassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var ChangePasswordReq Request.ChangePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&ChangePasswordReq)
	if err != nil {
		http.Error(w, "Gagal membaca body permintaan", http.StatusBadRequest)
		return
	}

	err = UserRepostory.ChangePassword(ctx, ChangePasswordReq)

	if err != nil {
		log.Fatal(err)
	}

	res := map[string]string{
		"status": "Change Password Successfully",
	}

	utilitis.ResponseJSON(w, res, http.StatusCreated)
}

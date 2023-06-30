package main

import (
	"context"
	"encoding/json"
	"final-project-go/Request"
	"final-project-go/middleware"
	"final-project-go/models"
	"final-project-go/repository/CarRepository"
	"final-project-go/repository/UserRepostory"
	"final-project-go/utilitis"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/change-password", middleware.JWTMiddleware(ChangePassword))
	router.GET("/list-user", middleware.JWTMiddleware(ListUser))
	router.POST("/cars", middleware.JWTMiddleware(CreateCar))
	http.ListenAndServe(":8080", router)
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

	token, err := middleware.CreateToken(user.NAME)
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

func CreateCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cencel := context.WithCancel(context.Background())
	defer cencel()

	var carReqeust Request.Car

	err := json.NewDecoder(r.Body).Decode(&carReqeust)
	fmt.Println(err)

	if err != nil {
		http.Error(w, "Gagal membaca body permintaan", http.StatusBadRequest)
		return
	}

	err = CarRepository.Insert(ctx, carReqeust)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := map[string]string{
		"status": "Change Password Successfully",
	}

	utilitis.ResponseJSON(w, res, http.StatusOK)
}

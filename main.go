package main

import (
	"context"
	"encoding/json"
	"final-project-go/Request"
	"final-project-go/docs"
	"final-project-go/middleware"
	"final-project-go/models"
	"final-project-go/repository/CarRepository"
	"final-project-go/repository/TourRepository"
	"final-project-go/repository/UserRepostory"
	"final-project-go/utilitis"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/change-password", middleware.JWTMiddleware(ChangePassword))
	router.GET("/list-user", middleware.JWTMiddleware(ListUser))
	router.POST("/cars", middleware.JWTMiddleware(CreateCar))
	router.GET("/cars", middleware.JWTMiddleware(listCar))
	router.GET("/cars/:id", middleware.JWTMiddleware(GetCar))
	router.PUT("/cars/:id", middleware.JWTMiddleware(UpdateCar))
	router.DELETE("/cars/:id", middleware.JWTMiddleware(DeleteCar))
	router.POST("/tours", middleware.JWTMiddleware(CreateTour))

	router.GET("/swagger/*filepath", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		filepath := ps.ByName("filepath")
		filepath = filepath[1:]

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		http.ServeFile(w, r, "docs/"+filepath)
	})
	//router.GET("/swagger/doc", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//	http.ServeFile(w, r, "docs/swagger.yaml")
	//})
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

// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.Users true "User details"
// @Success 201 {object} map[string]string
// @Router /register [post]
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

// @Summary Login
// @Description Authenticate user with provided credentials and generate access token
// @Tags User
// @Accept json
// @Produce json
// @Param request body Request.LoginRequest true "Login details"
// @Success 201 {object} map[string]string
// @Router /login [post]
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

func listCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx, cencel := context.WithCancel(context.Background())
	defer cencel()

	cars, err := CarRepository.GetAll(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utilitis.ResponseJSON(w, cars, http.StatusOK)
}

func GetCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ctx, cencel := context.WithCancel(context.Background())
	defer cencel()

	id := p.ByName("id")
	car, err := CarRepository.GetCar(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utilitis.ResponseJSON(w, car, http.StatusOK)
}

func UpdateCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
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

	err = CarRepository.Update(ctx, carReqeust, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utilitis.ResponseJSON(w, "Update Successfully", http.StatusOK)
}

func DeleteCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cencel := context.WithCancel(context.Background())
	defer cencel()

	err := CarRepository.Delete(ctx, id)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utilitis.ResponseJSON(w, "Deleted SuccessFully", http.StatusOK)
}

func CreateTour(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cencel := context.WithCancel(context.Background())
	defer cencel()

	var TouReq Request.TourRequest

	err := json.NewDecoder(r.Body).Decode(&TouReq)

	fmt.Println(err)

	if err != nil {
		http.Error(w, "form body harus di isi semua", http.StatusBadRequest)
		return
	}

	err = TourRepository.Insert(ctx, TouReq)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utilitis.ResponseJSON(w, "create tour successfuly", http.StatusCreated)
}

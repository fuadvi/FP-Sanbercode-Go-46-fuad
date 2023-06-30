package models

type Car struct {
	ID          int    `json:"id"`
	TITLE       string `json:"title"`
	PRICE       string `json:"price"`
	IMAGE       string `json:"image"`
	DESCRIPTION string `json:"description"`
	PASSENGER   int    `json:"passenger"`
	LUGGAGE     int    `json:"luggage"`
	CARTYPE     string `json:"car_type"`
	ISDRIVER    bool   `json:"is_driver"`
	DURATION    string `json:"duration"`
}

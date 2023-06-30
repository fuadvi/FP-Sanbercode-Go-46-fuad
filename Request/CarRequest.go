package Request

type Car struct {
	TITLE       string `json:"title"`
	PRICE       int    `json:"price"`
	IMAGE       string `json:"image"`
	DESCRIPTION string `json:"description"`
	PASSENGER   int    `json:"passenger"`
	LUGGAGE     int    `json:"luggage"`
	CARTYPE     string `json:"car_type"`
	ISDRIVER    bool   `json:"is_driver"`
	DURATION    string `json:"duration"`
}

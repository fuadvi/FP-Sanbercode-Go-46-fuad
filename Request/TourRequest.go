package Request

type TourRequest struct {
	TITLE       string `json:"title"`
	PRICE       string `json:"price"`
	DESCRIPTION string `json:"description"`
	DURATION    string `json:"duration"`
}

package models

type Tour struct {
	ID          int    `json:"id"`
	TITLE       string `json:"title"`
	PRICE       string `json:"price"`
	DESCRIPTION string `json:"description"`
	DURATION    string `json:"duration"`
}

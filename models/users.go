package models

type Users struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	EMAIL    string `json:"email"`
	NOHP     string `json:"no_hp"`
	PASSWORD string `json:"password"`
}

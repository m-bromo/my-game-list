package models

type GameOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	ImageUrl    string `json:"image_url"`
}

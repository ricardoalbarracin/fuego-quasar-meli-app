package model

type Response struct {
	Position Point  `json:"position"`
	Message  string `json:"message"`
}

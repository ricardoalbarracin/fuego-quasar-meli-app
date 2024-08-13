package model

type Satellites struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

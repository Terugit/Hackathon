package model

import "time"

type Thank struct {
	Id        string    `json:"id"`
	From_     string    `json:"from_"`
	To_       string    `json:"to_"`
	Point     int       `json:"point"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Thanks []Thank

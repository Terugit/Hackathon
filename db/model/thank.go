package model

type Thank struct {
	Id      string `json:"id"`
	From_   string `json:"from_"`
	To_     string `json:"to_"`
	Point   int    `json:"point"`
	Message string `json:"message"`
	// TimeStamp date? `json:"timestamp"`
}

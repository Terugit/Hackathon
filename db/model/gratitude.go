package model

type Gratitude struct {
	Id    string `json:"id"`
	From  string `json:"from"`
	To    string `json:"to"`
	Point int    `json:"point"`
	// TimeStamp date? `json:"timestamp"`
}

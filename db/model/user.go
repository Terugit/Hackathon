package model

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id" gorm:"not null"`
	Point     int    `json:"point"`
}

type Users []User

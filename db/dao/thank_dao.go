package dao

import (
	"db/model"
	"gorm.io/gorm"
)

func FindThank(p *model.Thank, contributionId string) (tx *gorm.DB) {
	return db.Where("id = ?", contributionId).First(&p)
}

func CreateThank(p *model.Thank) (tx *gorm.DB) {
	return db.Create(&p)
}

func DeleteThank(p *model.Thank, contributionId string) (tx *gorm.DB) {
	return db.Where("id = ?", contributionId).Delete(&p)
}

func UpdateThank(p *model.Thank, contributionId string, point int, message string) (tx *gorm.DB) {
	return db.Model(&p).Where("id = ?", contributionId).Updates(model.Thank{Message: message, Point: point})
}

func GetAllThank(u *model.Thanks) (tx *gorm.DB) {
	return db.Find(&u)
}

func GetAllThankSent(u *model.Thanks, from_ string) (tx *gorm.DB) {
	return db.Where(model.Thank{From_: from_}).Find(&u)
}

func GetAllThankReceived(u *model.Thanks, to_ string) (tx *gorm.DB) {
	return db.Where(model.Thank{To_: to_}).Find(&u)
}
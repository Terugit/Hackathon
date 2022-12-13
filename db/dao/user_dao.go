package dao

import (
	"db/model"
	"gorm.io/gorm"
)

func CreateUser(p *model.User) (tx *gorm.DB) {
	return db.Create(&p)
}

func DeleteUser(p *model.User, userId string) (tx *gorm.DB) {
	return db.Where("id = ?", userId).Delete(&p)
}

//func UpdateUserAttributes(p *model.User, userId string, userName string, description string, avatarUrl string) (tx *gorm.DB) {
//	return db.Model(&p).Where("id = ?", userId).Updates(model.User{Name: userName, Description: description, AvatarUrl: avatarUrl})
//}

func GetAllUsers(u *model.Users) (tx *gorm.DB) {
	return db.Find(&u)
}

func FindUserById(p *model.User, accountId string) (tx *gorm.DB) {
	return db.Where("workspace_id = ? AND account_id = ?", accountId).First(&p)
}

func FindUserByUserId(p *model.User, userId string) (tx *gorm.DB) {
	return db.Where("id = ?", userId).First(&p)
}

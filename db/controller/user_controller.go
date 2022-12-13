package controller

import (
	"db/dao"
	"db/model"
	"db/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserReq struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type UserRes struct {
	Id        string `json:"user_id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id"`
}
type UsersRes []UserRes

type UserInfoRes struct {
	Id        string `json:"user_id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id"`
}

func FetchUserInfo(c *gin.Context) {
	userId := utils.GetValueFromContext(c, "userId")

	targetUser := model.User{}
	if err := dao.FindUserByUserId(&targetUser, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &UserInfoRes{
		targetUser.Id,
		targetUser.Name,
		targetUser.AccountId,
	}

	c.JSON(http.StatusOK, res)
}

func CreateUser(c *gin.Context) {

	req := new(CreateUserReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//targetAccount := model.Account{}
	//if err := dao.FindAccountByEmail(&targetAccount, req.Email).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//if targetAccount.Id == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"message": "account not found"})
	//	return
	//}
	//
	//targetUser := model.User{
	//	Id: "",
	//}
	//err := dao.FindUserById(&targetUser, targetAccount.Id).Error
	//
	//if err != gorm.ErrRecordNotFound {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//if targetUser.Id != "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
	//	return
	//}

	newUserId := utils.GenerateId()
	newUser := model.User{
		Id:   newUserId,
		Name: req.Name,
		//AccountId:   targetAccount.Id,
	}

	if err := dao.CreateUser(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &UserRes{
		newUser.Id,
		newUser.Name,
		newUser.AccountId,
	}

	c.JSON(http.StatusOK, res)
}

func FetchAllUsers(c *gin.Context) {

	targetUsers := model.Users{}
	if err := dao.GetAllUsers(&targetUsers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make(UsersRes, 0)
	for i := 0; i < len(targetUsers); i++ {
		res = append(res, UserRes{
			targetUsers[i].Id,
			targetUsers[i].Name,
			targetUsers[i].AccountId,
		})
	}

	c.JSON(http.StatusOK, res)
}

func DeleteUser(c *gin.Context) {
	userId := utils.GetValueFromContext(c, "userId")
	role := utils.GetValueFromContext(c, "role")

	if role == "owner" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cant delete yourself"})
		return
	}

	targetUser := model.User{}
	if err := dao.DeleteUser(&targetUser, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

package controller

import (
	"db/dao"
	"db/model"
	"db/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const layout = "1997-07-01 21:07:20"

type CreateThankReq struct {
	ReceiverId string `json:"receiver_id" binding:"required"`
	Point      int    `json:"points" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

type DeleteThankReq struct {
	ThankId string `json:"thank_id" binding:"required"`
}

type EditThankReq struct {
	ThankId string `json:"thank_id" binding:"required"`
	Point   int    `json:"point" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type ThankRes struct {
	Id        string `json:"thank_id"`
	From_     string `json:"sender_id"`
	To_       string `json:"receiver_id"`
	Point     int    `json:"point"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type ThanksRes []ThankRes

type EditThankRes struct {
	Id       string `json:"thank_id"`
	To_      string `json:"receiver_id"`
	Point    int    `json:"point"`
	Message  string `json:"message"`
	UpdateAt string `json:"update_at"`
}

type ThankReport struct {
	UserId        string `json:"user_id"`
	Name          string `json:"name"`
	ThankSent     int    `json:"thank_sent"`
	PointSent     int    `json:"points_sent"`
	ThankReceived int    `json:"thank_received"`
	PointReceived int    `json:"points_received"`
}
type ThankReportRes []ThankReport

//func FetchThankReport(c *gin.Context) {
//
//	endDate := time.Now()
//	startDate := endDate.Add(-7 * 24 * time.Hour)
//
//	targetThanks := model.Thanks{}
//	if err := dao.GetDesignatedThankInWorkspace(&targetThanks, startDate, endDate).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	targetUsers := model.Users{}
//	if err := dao.GetAllUsersInWorkspace(&targetUsers, workspaceId).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	res := make(ThankReportRes, 0)
//
//	for i := 0; i < len(targetUsers); i++ {
//		thankSent := 0
//		pointsSent := 0
//		reactionSent := 0
//		thankReceived := 0
//		pointsReceived := 0
//		reactionReceived := 0
//		for j := 0; j < len(targetThanks); j++ {
//			if targetUsers[i].Id == targetThanks[j].From {
//				thankSent++
//				pointsSent += targetThanks[j].Point
//				reactionSent += targetThanks[j].Reaction
//			}
//
//			if targetUsers[i].Id == targetThanks[j].To {
//				thankReceived++
//				pointsReceived += targetThanks[j].Point
//				reactionReceived += targetThanks[j].Reaction
//			}
//		}
//		res = append(res, ThankReport{
//			targetUsers[i].Id,
//			targetUsers[i].Name,
//			thankSent,
//			pointsSent,
//			thankReceived,
//			pointsReceived,
//		})
//	}
//
//	c.JSON(http.StatusOK, res)
//}

func CreateThank(c *gin.Context) {
	req := new(CreateThankReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := utils.GetValueFromContext(c, "userId")

	if userId == req.ReceiverId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cannot send thank to yourself"})
		return
	}

	if req.Point < 1 && req.Point > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cannot send more than 100 points"})
		return
	}

	thankId := utils.GenerateId()
	newThank := model.Thank{
		Id:      thankId,
		From_:   userId,
		To_:     req.ReceiverId,
		Point:   req.Point,
		Message: req.Message,
	}

	if err := dao.CreateThank(&newThank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	res := &ThankRes{
		newThank.Id,
		newThank.From_,
		newThank.To_,
		newThank.Point,
		newThank.Message,
		newThank.CreatedAt.In(jst).Format(layout),
		newThank.UpdatedAt.In(jst).Format(layout),
	}

	c.JSON(http.StatusOK, res)
}

func DeleteThank(c *gin.Context) {
	req := new(DeleteThankReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	targetThank := model.Thank{}
	if err := dao.DeleteThank(&targetThank, req.ThankId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "thank deleted"})
}

func EditThank(c *gin.Context) {
	req := new(EditThankReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Point > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cant send more than 100 points"})
		return
	}

	targetThank := model.Thank{}
	if err := dao.UpdateThank(&targetThank, req.ThankId, req.Point, req.Message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	res := &EditThankRes{
		req.ThankId,
		targetThank.To_,
		targetThank.Point,
		targetThank.Message,
		targetThank.UpdatedAt.In(jst).Format(layout),
	}

	c.JSON(http.StatusOK, res)
}

func FetchAllThank(c *gin.Context) {

	targetThanks := model.Thanks{}
	if err := dao.GetAllThank(&targetThanks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if targetThanks[0].Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thanks not found"})
		return
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	res := make(ThanksRes, 0)
	for i := 0; i < len(targetThanks); i++ {
		res = append(res, ThankRes{
			targetThanks[i].Id,
			targetThanks[i].From_,
			targetThanks[i].To_,
			targetThanks[i].Point,
			targetThanks[i].Message,
			targetThanks[i].CreatedAt.In(jst).Format(layout),
			targetThanks[i].UpdatedAt.In(jst).Format(layout),
		})
	}

	c.JSON(http.StatusOK, res)
}

func FetchAllThankSent(c *gin.Context) {
	userId := utils.GetValueFromContext(c, "userId")

	targetThanks := model.Thanks{}
	if err := dao.GetAllThankSent(&targetThanks, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if targetThanks[0].Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thanks not found"})
		return
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	res := make(ThanksRes, 0)
	for i := 0; i < len(targetThanks); i++ {
		res = append(res, ThankRes{
			targetThanks[i].Id,
			targetThanks[i].From_,
			targetThanks[i].To_,
			targetThanks[i].Point,
			targetThanks[i].Message,
			targetThanks[i].CreatedAt.In(jst).Format(layout),
			targetThanks[i].UpdatedAt.In(jst).Format(layout),
		})
	}

	c.JSON(http.StatusOK, res)
}

func FetchAllThankReceived(c *gin.Context) {
	userId := utils.GetValueFromContext(c, "userId")

	targetThanks := model.Thanks{}
	if err := dao.GetAllThankReceived(&targetThanks, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if targetThanks[0].Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thanks not found"})
		return
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	res := make(ThanksRes, 0)
	for i := 0; i < len(targetThanks); i++ {
		res = append(res, ThankRes{
			targetThanks[i].Id,
			targetThanks[i].From_,
			targetThanks[i].To_,
			targetThanks[i].Point,
			targetThanks[i].Message,
			targetThanks[i].CreatedAt.In(jst).Format(layout),
			targetThanks[i].UpdatedAt.In(jst).Format(layout),
		})
	}

	c.JSON(http.StatusOK, res)
}

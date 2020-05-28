/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-02 18:06:20
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-09 11:50:51
 */

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
	"api.ys1994-user/utils"
)

type saveVodCommentValidator struct {
	VodID               uint                `json:"vodId" binding:"required,min=1"`
	Content             string              `json:"content" binding:"required,min=1,max=2000"`
	VodStar             uint8               `json:"vodStar" binding:"max=5"`
	DirectorStar        uint8               `json:"directorStar" binding:"max=5"`
	ActorStar           uint8               `json:"actorStar" binding:"max=5"`
	ActressStar         uint8               `json:"actressStar" binding:"max=5"`
	SuppActorStar       uint8               `json:"suppActorStar" binding:"max=5"`
	SuppActressStar     uint8               `json:"suppActressStar" binding:"max=5"`
	ScreenplayStar      uint8               `json:"screenplayStar" binding:"max=5"`
	CinematographyStar  uint8               `json:"cinematographyStar" binding:"max=5"`
	EditStar            uint8               `json:"editStar" binding:"max=5"`
	SoundStar           uint8               `json:"soundStar" binding:"max=5"`
	VisualStar          uint8               `json:"visualStar" binding:"max=5"`
	MakeupStar          uint8               `json:"makeupStar" binding:"max=5"`
	CostumeStar         uint8               `json:"costumeStar" binding:"max=5"`
	MusicStar           uint8               `json:"musicStar" binding:"max=5"`
	AtUserIDs           []uint              `json:"atUserIds" binding:"max=10"`
}

// SaveVodCommentHandle 保存影片评论
func SaveVodCommentHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData saveVodCommentValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	vodID := hasVod(formData.VodID)
	if vodID == 0 {
		cJSONNotFound(c, "影片不存在")
		return
	}

	var uvc model.UserVodComment
	uvc.Super()
	uvc.Filter.UserID = userID
	uvc.Filter.VodID = vodID
	uvc.Detail()

	if err := utils.Copy(formData, &uvc.Filter); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	uvc.Filter.UserID = userID
	uvc.Filter.Grade = getUserGrade(userID)

	if uvc.ID == 0 {
		uvc.Create()
	} else {
		uvc.Filter.ID = uvc.ID
		uvc.Update()
	}

	go vodCommentAtHandle(formData.AtUserIDs, uvc.Filter.ID)

	cJSONOk(c, uvc.Filter.ID)
}

func vodCommentAtHandle(userIDs []uint, vodCommentID uint) {
	for _, userID := range userIDs {
		var vca model.UserVodCommentAt
		vca.Super()
		vca.Filter.UserID = userID
		vca.Filter.UserVodCommentID = vodCommentID
		vca.Detail()
		vca.Filter.Read = 1
		if vca.ID == 0 {
			vca.Create()
		} else {
			vca.Filter.ID = vca.ID
			vca.Update()
		}
	}
}

// GetVodCommentListHandle 影片评论列表
func GetVodCommentListHandle(c *gin.Context) {
	var err error
	var vcl model.UserVodCommentList
	vcl.Page, vcl.Size, err = listHandle(c, &vcl.Filter, &vcl.Filters, &vcl.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
	} else {
		vcl.FilterUserLikeVodComment.UserID, _ = getUserID(c)
		vcl.Find()
		cJSONOk(c, vcl)
	}
}

// GetVodCommentDetailHandle 影片评论
func GetVodCommentDetailHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		cJSONBadRequest(c, "id 不能为空")
		return
	}

	errStr, uvc := hasVodComment(uint(id));
	if errStr != "" {
		cJSONBadRequest(c, errStr)
		return
	}

	userID, _ := getUserID(c)
	if userID > 0 {
		uvc.UserLikeVodComment.Super()
		uvc.UserLikeVodComment.Filter.UserVodCommentID = uvc.ID
		uvc.UserLikeVodComment.Filter.UserID = userID
		uvc.UserLikeVodComment.Detail()
	}

	cJSONOk(c, uvc)
}

func hasVodComment(id uint) (string, *model.UserVodComment) {
	var uvc model.UserVodComment
	uvc.Super()
	uvc.Filter.ID = id
	uvc.Base.Detail()
	if uvc.ID == 0 {
		return "评论不存在", &uvc
	}

	return "", &uvc
}

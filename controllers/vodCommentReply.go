/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-02 18:06:20
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-09 11:49:46
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type saveVodCommentReplyValidator struct {
	VodCommentID        uint                `json:"vodCommentId" binding:"required,min=1"`
	Content             string              `json:"content" binding:"required,min=1,max=200"`
	AtUserIDs           []uint              `json:"atUserIds" binding:"max=10"`
}

// SaveVodCommentReplyHandle 保存影片评论回复
func SaveVodCommentReplyHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData saveVodCommentReplyValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	errStr, uvc := hasVodComment(formData.VodCommentID);
	if errStr != "" {
		cJSONBadRequest(c, errStr)
		return
	}

	var vcr model.UserVodCommentReply
	vcr.Super()
	vcr.Filter.UserID = userID
	vcr.Filter.UserVodCommentID = formData.VodCommentID
	vcr.Filter.Content = formData.Content
	vcr.Filter.Grade = getUserGrade(userID)
	vcr.Create()

	uvc.IncrementReply()

	if uvc.UserID != userID {
		formData.AtUserIDs = append(formData.AtUserIDs, uvc.UserID)
	}
	go vodCommentReplyAtHandle(formData.AtUserIDs, vcr.Filter.ID)

	cJSONOk(c, vcr.Filter.ID)
}

func vodCommentReplyAtHandle(userIDs []uint, vodCommentReplyID uint) {
	for _, userID := range userIDs {
		var vca model.UserVodCommentReplyAt
		vca.Super()
		vca.Filter.UserID = userID
		vca.Filter.UserVodCommentReplyID = vodCommentReplyID
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

// GetVodCommentReplyListHandle 影片评论回复列表
func GetVodCommentReplyListHandle(c *gin.Context) {
	var err error
	var vcl model.UserVodCommentReplyList
	vcl.Page, vcl.Size, err = listHandle(c, &vcl.Filter, &vcl.Filters, &vcl.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
	} else {
		vcl.FilterUserLikeVodCommentReply.UserID, _ = getUserID(c)
		vcl.Find()
		cJSONOk(c, vcl)
	}
}

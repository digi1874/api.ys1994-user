/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-02 18:06:20
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-10 16:04:41
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type saveVodCollectionCommentValidator struct {
	VodCollectionID     uint                `json:"vodCollectionId" binding:"required,min=1"`
	Content             string              `json:"content" binding:"required,min=1,max=200"`
	AtUserIDs           []uint              `json:"atUserIds" binding:"max=10"`
}

// SaveVodCollectionCommentHandle 保存影片评论回复
func SaveVodCollectionCommentHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData saveVodCollectionCommentValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	errStr, vc := hasVodCollection(formData.VodCollectionID, 0)
	if errStr != "" {
		cJSONBadRequest(c, errStr)
		return
	}

	var vcc model.UserVodCollectionComment
	vcc.Super()
	vcc.Filter.UserID = userID
	vcc.Filter.UserVodCollectionID = formData.VodCollectionID
	vcc.Filter.Content = formData.Content
	vcc.Filter.Grade = getUserGrade(userID)
	vcc.Create()

	if vcc.UserID != userID {
		formData.AtUserIDs = append(formData.AtUserIDs, vc.UserID)
	}
	go vodCollectionCommentAtHandle(formData.AtUserIDs, vcc.Filter.ID)

	cJSONOk(c, vcc.Filter.ID)
}

func vodCollectionCommentAtHandle(userIDs []uint, vodCollectionCommentID uint) {
	for _, userID := range userIDs {
		var vca model.UserVodCollectionCommentAt
		vca.Super()
		vca.Filter.UserID = userID
		vca.Filter.UserVodCollectionCommentID = vodCollectionCommentID
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

// GetVodCollectionCommentListHandle 影片评论回复列表
func GetVodCollectionCommentListHandle(c *gin.Context) {
	var err error
	var vcl model.UserVodCollectionCommentList
	vcl.Page, vcl.Size, err = listHandle(c, &vcl.Filter, &vcl.Filters, &vcl.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
	} else {
		vcl.FilterUserLikeVodCollectionComment.UserID, _ = getUserID(c)
		vcl.Find()
		cJSONOk(c, vcl)
	}
}

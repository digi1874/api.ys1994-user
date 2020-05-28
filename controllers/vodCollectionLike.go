/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-27 23:13:38
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-09 13:19:11
 */

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type likeVodCollectionValidator struct {
	VodCollectionID uint `json:"vodCollectionId" binding:"required,min=1"`
}

// IsLikeVodCollectionHandle is like vod Collection
func IsLikeVodCollectionHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	vodCollectionID, _ := strconv.Atoi(c.Param("id"))
	if vodCollectionID < 1 {
		cJSONBadRequest(c, "集合 id 不能为空")
		return
	}

	var lvc model.UserLikeVodCollection
	lvc.Super()
	lvc.Filter.UserID = userID
	lvc.Filter.UserVodCollectionID = uint(vodCollectionID)
	lvc.Filter.State = 1
	lvc.Detail()
	if lvc.ID == 0 {
		cJSONOk(c, false)
	} else {
		cJSONOk(c, true)
	}
}

// LikeVodCollectionHandle Like Vod Collection
func LikeVodCollectionHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData likeVodCollectionValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	if errStr, _ := hasVodCollection(formData.VodCollectionID, 0); errStr != "" {
		cJSONBadRequest(c, errStr)
		return
	}

	var lvc model.UserLikeVodCollection
	lvc.Super()
	lvc.Filter.UserID = userID
	lvc.Filter.UserVodCollectionID = formData.VodCollectionID
	lvc.Detail()

	if lvc.ID == 0 {
		lvc.Create()
		go vodCollectionIncrementLike(lvc.Filter.UserVodCollectionID)
	} else {
		if lvc.State == 1 {
			lvc.Filter.State = 2
			go vodCollectionDecrementLike(lvc.Filter.UserVodCollectionID)
		} else {
			lvc.Filter.State = 1
			go vodCollectionIncrementLike(lvc.Filter.UserVodCollectionID)
		}
		lvc.Filter.ID = lvc.ID
		lvc.Update()
	}

	cJSONOk(c, lvc.Filter.State)
}

// GetVodCollectionLikeListHandle 收藏视频列表
func GetVodCollectionLikeListHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	var err error
	var vcl model.UserLikeVodCollectionList
	vcl.Page, vcl.Size, err = listHandle(c, &vcl.Filter, &vcl.Filters, &vcl.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	vcl.FilterUserLikeVodCollection.UserID = userID
	vcl.FilterUserLikeVodCollection.State = 1
	vcl.Find()

	cJSONOk(c, vcl)
}

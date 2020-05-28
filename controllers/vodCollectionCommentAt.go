/*
 * @Author: lin.zhenhui
 * @Date: 2020-05-13 16:29:09
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-13 20:14:35
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type vodCollectionCommentAtIDsValidator struct {
	IDs []uint `json:"ids"`
}

//UpdateVodCollectionCommentAtReadHandle UpdateVodCollectionCommentAtReadHandle
func UpdateVodCollectionCommentAtReadHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData vodCollectionCommentAtIDsValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	var vca model.UserVodCollectionCommentAt
	vca.Filter.UserID = userID
	vca.HaveRead(formData.IDs)
	cJSONOk(c, true)
}

// DeleteVodCollectionCommentAtHandle 删除
func DeleteVodCollectionCommentAtHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData deleteValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	var vca model.UserVodCollectionCommentAt
	vca.Super()
	vca.Filter.UserID = userID
	vca.Delete(formData.IDs)
	cJSONOk(c, "删除成功")
}

// GetVodCollectionCommentAtListHandle 影片评论回复@列表
func GetVodCollectionCommentAtListHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}
	var err error
	var cal model.UserVodCollectionCommentAtList
	cal.Page, cal.Size, err = listHandle(c, &cal.Filter, &cal.Filters, &cal.Orders)

	if err != nil {
		cJSONBadRequest(c, err.Error())
	} else {
		cal.Filter.UserID = userID
		cal.Find()
		cJSONOk(c, cal)
	}
}

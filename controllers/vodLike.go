/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-27 23:13:38
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 21:25:39
 */

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type likeVodValidator struct {
	VodID uint `json:"vodId" binding:"required,min=1"`
}

// IsLikeVodHandle is like vod
func IsLikeVodHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	vID, _ := strconv.Atoi(c.Param("id"))
	vodID := hasVod(uint(vID))
	if vodID == 0 {
		cJSONNotFound(c, "影片不存在")
		return
	}

	var ulv model.UserLikeVod
	ulv.Super()
	ulv.Filter.UserID = userID
	ulv.Filter.VodID = vodID
	ulv.Filter.State = 1
	ulv.Detail()
	if ulv.ID == 0 {
		cJSONOk(c, false)
	} else {
		cJSONOk(c, true)
	}
}

// LikeVodHandle Like VodHandle
func LikeVodHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData likeVodValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	vodID := hasVod(formData.VodID)
	if vodID == 0 {
		cJSONNotFound(c, "影片不存在")
		return
	}

	var vodLike model.UserLikeVod
	vodLike.Super()
	vodLike.Filter.UserID = userID
	vodLike.Filter.VodID = vodID
	vodLike.Detail()

	if vodLike.ID == 0 {
		vodLike.Create()
		go vodDataIncrementLike(vodID)
	} else {
		if vodLike.State == 1 {
			vodLike.Filter.State = 2
			go vodDataDecrementLike(vodID)
		} else {
			vodLike.Filter.State = 1
			go vodDataIncrementLike(vodID)
		}
		vodLike.Filter.ID = vodLike.ID
		vodLike.Update()
	}

	cJSONOk(c, vodLike.Filter.State)
}

// GetVodLikeListHandle 收藏视频列表
func GetVodLikeListHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	var err error
	var vll model.UserLikeVodList
	vll.Page, vll.Size, err = listHandle(c, &vll.FilterVod, &vll.Filters, &vll.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	vll.Filter.UserID = userID
	vll.Filter.State = 1
	vll.Find()

	cJSONOk(c, vll)
}

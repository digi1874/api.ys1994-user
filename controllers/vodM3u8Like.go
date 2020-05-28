/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-27 23:13:38
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 21:25:39
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type likeVodM3u8Validator struct {
	VodM3u8ID uint `json:"vodM3u8Id" binding:"required,min=1"`
}

// LikeVodM3u8Handle Like VodM3u8Handle
func LikeVodM3u8Handle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData likeVodM3u8Validator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	vodM3u8ID := hasVodM3u8(formData.VodM3u8ID)
	if vodM3u8ID == 0 {
		cJSONNotFound(c, "视频不存在")
		return
	}

	var lvm model.UserLikeVodM3u8
	lvm.Super()
	lvm.Filter.UserID = userID
	lvm.Filter.VodM3u8ID = vodM3u8ID
	lvm.Detail()

	if lvm.ID == 0 {
		lvm.Create()
		go vodM3u8DataIncrementLike(vodM3u8ID)
	} else {
		if lvm.State == 1 {
			lvm.Filter.State = 2
			go vodM3u8DataDecrementLike(vodM3u8ID)
		} else {
			lvm.Filter.State = 1
			go vodM3u8DataIncrementLike(vodM3u8ID)
		}
		lvm.Filter.ID = lvm.ID
		lvm.Update()
	}

	cJSONOk(c, lvm.Filter.State)
}

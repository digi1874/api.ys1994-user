/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-28 16:37:51
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 19:27:19
 */

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type saveVodPlayValidator struct {
	VodM3u8ID uint `json:"m3u8Id" binding:"required,min=1"`
	Time      uint `json:"time" binding:"required,min=1"`
}

// SaveVodPlayHandle 保存播放记录
func SaveVodPlayHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData saveVodPlayValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	var m3u8 model.VodM3u8
	m3u8.Super()
	m3u8.Filter.ID = formData.VodM3u8ID
	m3u8.Detail()

	if m3u8.ID == 0 {
		cJSONBadRequest(c, "视频不存在")
		return
	}

	var upm model.UserPlayM3u8
	upm.Super()
	upm.Filter.UserID = userID
	upm.Filter.VodM3u8ID = m3u8.ID
	upm.Detail()

	upm.Filter.VodID = m3u8.VodID
	upm.Filter.Time = formData.Time
	if upm.ID == 0 {
		upm.Create()
	} else {
		upm.Filter.ID = upm.ID
		upm.Update()
	}

	go countUserPlayTime(userID)

	cJSONOk(c, true)
}

// DeleteVodPlayHandle 删除播放记录
func DeleteVodPlayHandle(c *gin.Context) {
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

	var upm model.UserPlayM3u8
	upm.Super()
	upm.Filter.UserID = userID
	upm.Delete(formData.IDs)
	cJSONOk(c, "删除成功")
}

// GetVodPlayDetailHandle Get Vod Play Detail
func GetVodPlayDetailHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	vodID, _ := strconv.Atoi(c.Param("vodId"))
	if vodID < 1 {
		cJSONBadRequest(c, "vodID不能为空")
		return
	}

	var upm model.UserPlayM3u8
	upm.Super()
	upm.Filter.UserID = userID
	upm.Filter.VodID = uint(vodID)
	upm.Detail()

	cJSONOk(c, upm)
}

// GetVodPlayListHandle Get Vod Play list
func GetVodPlayListHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	var err error
	var pml model.UserPlayM3u8List
	pml.Page, pml.Size, err = listHandle(c, &pml.FilterVod, &pml.Filters, &pml.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	pml.Filter.UserID = userID
	pml.Find()

	cJSONOk(c, pml)
}

// GetVodPlayM3u8TimeHandle Get Play M3u8 Time
func GetVodPlayM3u8TimeHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	m3u8ID, _ := strconv.Atoi(c.Param("m3u8Id"))
	if m3u8ID < 1 {
		cJSONBadRequest(c, "m3u8ID不能为空")
		return
	}

	var upm model.UserPlayM3u8
	upm.Super()
	upm.Filter.UserID = userID
	upm.Filter.VodM3u8ID = uint(m3u8ID)
	upm.Base.Detail()

	cJSONOk(c, upm)
}
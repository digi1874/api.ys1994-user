/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-07 18:42:49
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 21:02:38
 */

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

// GetUserVodM3u8DataHandle 一些用户与视频互动数据
func GetUserVodM3u8DataHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("vodM3u8Id"))
	vodM3u8ID := hasVodM3u8(uint(id))
	if vodM3u8ID == 0 {
		cJSONNotFound(c, "视频不存在")
		return
	}

	userID, _ := getUserID(c)

	var vmd model.UserVodM3u8Data
	vmd.VodM3u8Data.Filter.VodM3u8ID = vodM3u8ID
	vmd.FilterUser.ID = userID
	vmd.Detail()

	cJSONOk(c, vmd)
}

func vodM3u8DataIncrementLike(vodM3u8ID uint) {
	var vd model.VodM3u8Data
	vd.Super()
	vd.Filter.VodM3u8ID = vodM3u8ID
	vd.FirstOrCreate()
	vd.IncrementLike()
}

func vodM3u8DataDecrementLike(vodM3u8ID uint) {
	var vd model.VodM3u8Data
	vd.Super()
	vd.Filter.VodM3u8ID = vodM3u8ID
	vd.FirstOrCreate()
	vd.DecrementLike()
}

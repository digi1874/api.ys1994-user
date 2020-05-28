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

// GetUserVodDataHandle 一些用户与影片互动数据
func GetUserVodDataHandle(c *gin.Context) {
	vID, _ := strconv.Atoi(c.Param("vodId"))
	vodID := hasVod(uint(vID))
	if vodID == 0 {
		cJSONNotFound(c, "影片不存在")
		return
	}

	userID, _ := getUserID(c)

	var uvd model.UserVodData
	uvd.VodData.Filter.VodID = vodID
	uvd.FilterUser.ID = userID
	uvd.Detail()

	cJSONOk(c, uvd)
}

func vodDataIncrementLike(vodID uint) {
	var vd model.VodData
	vd.Super()
	vd.Filter.VodID = vodID
	vd.FirstOrCreate()
	vd.IncrementLike()
}

func vodDataDecrementLike(vodID uint) {
	var vd model.VodData
	vd.Super()
	vd.Filter.VodID = vodID
	vd.FirstOrCreate()
	vd.DecrementLike()
}

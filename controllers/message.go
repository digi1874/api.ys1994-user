/*
 * @Author: lin.zhenhui
 * @Date: 2020-05-13 15:34:04
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-13 15:55:54
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

// GetMessageUnreadCount GetMessageUnreadCount
func GetMessageUnreadCount(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	mc := model.MessageCount{ UserID: userID }
	mc.UnreadCount()

	cJSONOk(c, mc)
}

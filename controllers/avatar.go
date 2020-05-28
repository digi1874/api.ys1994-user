/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-17 15:38:48
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-03 18:03:21
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

// GetAvatarListHandle 头像列表
func GetAvatarListHandle(c *gin.Context) {
	var err error
	var ual model.UserAvatarList
	ual.Page, ual.Size, err = listHandle(c, &ual.Filter, &ual.Filters, &ual.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	ual.Filter.State = 1
	ual.Find()

	cJSONOk(c, ual)
}

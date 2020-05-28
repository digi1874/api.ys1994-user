/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-15 17:31:04
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-30 18:58:27
 */

package controllers


import (
	"time"

	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

// SearchUserHandle 查找user
func SearchUserHandle(c *gin.Context) {
	var err error
	var ul model.UserList
	ul.Page, ul.Size, err = listHandle(c, &ul.Filter, &ul.Filters, &ul.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}
	ul.Find()
	cJSONOk(c, ul)
}

func getUserID(c *gin.Context) (id uint, errStr string) {
	signature := c.GetHeader("auth")

	// 是否有signature
	if signature == "" {
		return 0, "未授权"
	}

	var userLogin model.UserLogin
	userLogin.Super()
	userLogin.Filter.Signature = signature
	userLogin.Detail()

	if userLogin.ID == 0 || userLogin.State != 1 || userLogin.UserID == 0 || userLogin.IP != c.ClientIP() {
		return 0, "登录无效，请重新登录"
	}

	if uint(time.Now().Unix()) > userLogin.Exp {
		return 0, "登录过期，请重新登录"
	}

	return userLogin.UserID, ""
}

func countUserPlayTime(userID uint) {
	var user model.User
	user.Super()
	user.Filter.ID = userID
	user.CountPlayTime()
}

func getUserGrade(id uint) (grade uint8) {
	var user model.User
	user.Super()
	user.Filter.ID = id
	user.Base.Detail()

	if user.PlayTime < 57600 {
		return 1
	}
	if user.PlayTime < 230400 {
		return 2
	}
	if user.PlayTime < 921600 {
		return 3
	}
	if user.PlayTime < 3686400 {
		return 4
	}
	return 5
}

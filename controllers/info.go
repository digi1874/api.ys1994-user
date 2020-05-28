/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-15 15:58:35
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-28 16:57:13
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
	"api.ys1994-user/utils"
)

type updateUserValidator struct {
	Avatar          string         `json:"avatar"`
	Name            string         `json:"name"`
	Sex             uint8          `json:"sex"`
	Birthday        uint           `json:"birthday"`
	Autograph       string         `json:"autograph"`
}

// GetInfoHandle 获取个人信息
func GetInfoHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	var user model.User
	user.Super()
	user.Filter.ID = userID
	user.Detail()

	if user.ID == 0 {
		cJSONBadRequest(c, "获取个人信息错误")
		return
	}

	if user.State != 1 {
		cJSONBadRequest(c, "账号已禁用")
		return
	}

	cJSONOk(c, user)
}

// UpdateInfoHandle 更新个人信息
func UpdateInfoHandle(c *gin.Context) {
	// user id
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData updateUserValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	var user model.User
	user.Super()
	// user filter
	if err := utils.Copy(formData, &user.Filter); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	user.Filter.ID = userID

	if formData.Name != "" && user.NameExist() {
		cJSONBadRequest(c, "昵称不可用")
		return
	}

	if formData.Avatar != "" {
		user.Filter.UserAvatarID = saveAvatar(userID, formData.Avatar)
	}

	user.Update()

	cJSONOk(c, "修改成功")
}

func saveAvatar(userID uint, image string) (avatarID uint) {
	var avatar model.UserAvatar
	avatar.Super()
	avatar.Filter.Image = image
	avatar.Detail()

	if avatar.ID == 0 {
		avatar.Filter.UserID = userID
		avatar.Create()
		return avatar.Filter.ID
	}

	avatar.Filter.ID = avatar.ID
	avatar.IncrementNum()
	return avatar.ID
}

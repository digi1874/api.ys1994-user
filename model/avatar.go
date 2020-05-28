/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-16 22:21:09
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 18:45:15
 */

package model

import (
	"api.ys1994-user/database"
)

// UserAvatar 用户头像
type UserAvatar struct {
	Base
	ID              uint                 `json:"id"`
	Image           string               `json:"image"`
	Filter          database.UserAvatar  `json:"-"`
}

// UserAvatarList 用户头像列表
type UserAvatarList struct {
	List
	Data            []UserAvatar         `json:"data"`
	Filter          database.UserAvatar  `json:"-"`
}

// Super Super
func (ua *UserAvatar) Super() {
	ua.Base.Super(ua, &ua.Filter)
}

// IncrementNum 增加数量
func (ua *UserAvatar) IncrementNum() {
	ua.Base.Increment("use_num")
}

// Find 列表
func (ua *UserAvatarList) Find() {
	// 不能查删除的
	ua.Filter.DeletedAt = nil
	db := database.DB.Where(&ua.Filter)
	ua.dbHandle(db.Order("use_num desc").Order("updated_time desc")).Find(&ua.Data)
	db.Model(&database.UserAvatar{}).Count(&ua.Count)
}

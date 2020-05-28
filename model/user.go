/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-15 16:55:08
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-13 19:26:50
 */

package model

import (
	"api.ys1994-user/database"
)

type user struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name"`
}

type userList struct {
	List
	Data         []user          `json:"data"`
	Filter       database.User   `json:"-"`
}

// USER USER
type USER struct {
	Base
	ID           uint           `json:"id"`
	Name         string         `json:"name"`
	Sex          uint8          `json:"sex"`
	UserAvatarID uint           `json:"-"`
	UserAvatar   UserAvatar     `json:"avatar"`
	PlayTime     uint           `json:"playTime"`
	Filter       database.User  `json:"-"`
}

// UserList User List
type UserList struct {
	List
	Data         []USER         `json:"data"`
	Filter       database.User  `json:"-"`
}

// User 用户
type User struct {
	USER
	Birthday        int64          `json:"birthday"`
	Autograph       string         `json:"autograph"`
	State           uint8          `json:"-"`
	Filter          database.User  `json:"-"`
}

// Super Super
func (u *User) Super() {
	u.Base.Super(u, &u.Filter)
}

// Detail 详情
func (u *User) Detail() {
	u.Base.Detail()
	if u.UserAvatarID != 0 {
		u.UserAvatar.Super()
		u.UserAvatar.Filter.ID = u.UserAvatarID
		u.UserAvatar.Filter.State = 1
		u.UserAvatar.Detail()
	}
}

// NameExist 检查昵称是否存在; u.Filter.ID  u.Filter.Name
func (u *User) NameExist() bool {
	userInfo := user{}
	database.DB.Where("id <> ?", u.Filter.ID).Where("`name` = ? OR `id` = ?", u.Filter.Name, u.Filter.Name).First(&userInfo)
	if userInfo.ID == 0 {
		return false
	}
	return true
}

// CountPlayTime 统计播放时间
func (u *User) CountPlayTime() {
	database.DB.Table("user_play_m3u8").Where("user_id = ?", u.Filter.ID).Select("sum(time)").Row().Scan(&u.Filter.PlayTime)
	u.Update()
}

// Find 列表
func (ul *UserList) Find() {
	db := ul.Filters.dbHandle(database.DB, "users")
	if ul.Filter.Name != "" {
		db = db.Where("`name` LIKE ?", "%" + ul.Filter.Name + "%")
		ul.Filter.Name = ""
	}
	db = db.Where(&ul.Filter)

	// 列表
	ul.dbHandle(db.Order("updated_time desc")).Preload("UserAvatar").Find(&ul.Data)
	// 总数
	db.Model(&database.User{}).Count(&ul.Count)
}

// Find 列表
func (ul *userList) Find() {
	db := ul.Filters.dbHandle(database.DB, "users")
	ul.dbHandle(db.Where(&ul.Filter)).Find(&ul.Data)
}

/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-30 21:55:49
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-09 13:14:43
 */

package model

import (
	"api.ys1994-user/database"
)

type userVodCollection struct {
	ID              uint                       `json:"id"`
	Name            string                     `json:"name"`
}

// UserVodCollection 视频集合
type UserVodCollection struct {
	Base
	ID              uint                       `json:"id"`
	UserID          uint                       `json:"userId"`
	Name            string                     `json:"name"`
	Pic             string                     `json:"pic"`
	Content         string                     `json:"content"`
	Share           uint8                      `json:"share"`
	State           uint8                      `json:"state"`
	LikeCount       uint                       `json:"likeCount"`
	Filter          database.UserVodCollection `json:"-"`
}

// UserVodCollectionList 视频集合列表
type UserVodCollectionList struct {
	List
	Data            []UserVodCollection         `json:"data"`
	Filter          database.UserVodCollection  `json:"-"`
}

// Super Super
func (uvc *UserVodCollection) Super() {
	uvc.Base.Super(uvc, &uvc.Filter)
}

// IncrementLike Increment Like
func (uvc *UserVodCollection) IncrementLike() {
	uvc.Increment("like_count")
}

// DecrementLike Decrement Like
func (uvc *UserVodCollection) DecrementLike() {
	uvc.Decrement("like_count")
}

// Find 列表
func (vcl *UserVodCollectionList) Find() {
	db := database.DB

	if vcl.Filter.Name != "" {
		db = db.Where("`name` LIKE ?", "%" + vcl.Filter.Name + "%")
		vcl.Filter.Name = ""
	}

	// 不能查删除的
	vcl.Filter.DeletedAt = nil
	db = db.Where(&vcl.Filter)

	vcl.dbHandle(db.Order("`id` desc")).Find(&vcl.Data)

	db.Model(&database.UserVodCollection{}).Count(&vcl.Count)
}

/*
 * @Author: lin.zhenhui
 * @Date: 2020-05-10 15:46:42
 * @Last Modified by:   lin.zhenhui
 * @Last Modified time: 2020-05-10 15:46:42
 */

package model

import (
	"api.ys1994-user/database"
)

// UserVodCollectionCommentAt 影片集合评论@
type UserVodCollectionCommentAt struct {
	Base
	ID                         uint                                `json:"id"`
	UserVodCollectionCommentID uint                                `json:"-"`
	UserVodCollectionComment   userVodCollectionComment            `json:"vodCollectionComment"`
	Read                       uint8                               `json:"read"`
	UpdatedTime                uint                                `json:"updatedTime"`
	Filter                     database.UserVodCollectionCommentAt `json:"-"`
}

// UserVodCollectionCommentAtList Vod Comment Reply At List
type UserVodCollectionCommentAtList struct {
	List
	Data            []UserVodCollectionCommentAt        `json:"data"`
	Filter          database.UserVodCollectionCommentAt `json:"-"`
}

// Super Super
func (vca *UserVodCollectionCommentAt) Super() {
	vca.Base.Super(vca, &vca.Filter)
}

// HaveRead HaveRead
func (vca *UserVodCollectionCommentAt) HaveRead(ids []uint) {
	db := database.DB.Model(&database.UserVodCollectionCommentAt{}).Where(&vca.Filter)
	if len(ids) > 0 {
		db = db.Where("id IN (?)", ids)
	}
	db.UpdateColumn("read", 2)
}

// Find 列表
func (cal *UserVodCollectionCommentAtList) Find() {
	// 不能查删除的
	cal.Filter.DeletedAt = nil
	db := database.DB.Where(&cal.Filter)

	cal.dbHandle(db).Preload("UserVodCollectionComment").Find(&cal.Data)

	db.Model(&database.UserVodCollectionCommentAt{}).Count(&cal.Count)

	if len(cal.Data) == 0 {
		return
	}

	var users map[uint]*user
	users = make(map[uint]*user)
	var ul userList

	for index, vca := range cal.Data {
		if users[vca.UserVodCollectionComment.UserID] == nil {
			users[vca.UserVodCollectionComment.UserID] = &user{}
		}
		cal.Data[index].UserVodCollectionComment.User = users[vca.UserVodCollectionComment.UserID]
		ul.Filters.IDs = append(ul.Filters.IDs, vca.UserVodCollectionComment.UserID)
	}

	ul.Find()
	for _, user := range ul.Data {
		users[user.ID].ID = user.ID
		users[user.ID].Name = user.Name
	}
}

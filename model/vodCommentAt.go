/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-02 20:37:48
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-13 20:15:59
 */

package model

import (
	"api.ys1994-user/database"
)

// UserVodCommentAt 影片评论@
type UserVodCommentAt struct {
	Base
	ID               uint                      `json:"id"`
	UserVodCommentID uint                      `json:"-"`
	UserVodComment   userVodComment            `json:"vodComment"`
	Read             uint8                     `json:"read"`
	UpdatedTime      uint                      `json:"updatedTime"`
	Filter           database.UserVodCommentAt `json:"-"`
}

// UserVodCommentAtList Vod Comment At List
type UserVodCommentAtList struct {
	List
	Data            []UserVodCommentAt             `json:"data"`
	Filter          database.UserVodCommentAt      `json:"-"`
}

// Super Super
func (vca *UserVodCommentAt) Super() {
	vca.Base.Super(vca, &vca.Filter)
}

// HaveRead HaveRead
func (vca *UserVodCommentAt) HaveRead(ids []uint) {
	db := database.DB.Model(&database.UserVodCommentAt{}).Where(&vca.Filter)
	if len(ids) > 0 {
		db = db.Where("id IN (?)", ids)
	}
	db.UpdateColumn("read", 2)
}

// Find 列表
func (cal *UserVodCommentAtList) Find() {
	// 不能查删除的
	cal.Filter.DeletedAt = nil
	db := database.DB.Where(&cal.Filter)

	cal.dbHandle(db).Preload("UserVodComment").Find(&cal.Data)

	db.Model(&database.UserVodCommentAt{}).Count(&cal.Count)

	if len(cal.Data) == 0 {
		return
	}

	var users map[uint]*user
	users = make(map[uint]*user)
	var ul userList

	var vods map[uint]*vod
	vods = make(map[uint]*vod)
	var vl vodList

	for index, vca := range cal.Data {
		if users[vca.UserVodComment.UserID] == nil {
			users[vca.UserVodComment.UserID] = &user{}
		}
		cal.Data[index].UserVodComment.User = users[vca.UserVodComment.UserID]
		ul.Filters.IDs = append(ul.Filters.IDs, vca.UserVodComment.UserID)

		if vods[vca.UserVodComment.VodID] == nil {
			vods[vca.UserVodComment.VodID] = &vod{}
		}
		cal.Data[index].UserVodComment.Vod = vods[vca.UserVodComment.VodID]
		vl.Filters.IDs = append(vl.Filters.IDs, vca.UserVodComment.VodID)
	}

	ul.Find()
	for _, user := range ul.Data {
		users[user.ID].ID = user.ID
		users[user.ID].Name = user.Name
	}

	vl.Find()
	for _, vod := range vl.Data {
		vods[vod.ID].ID = vod.ID
		vods[vod.ID].Name = vod.Name
	}
}

/*
 * @Author: lin.zhenhui
 * @Date: 2020-05-09 11:43:24
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-09 11:45:08
 */

package model

import (
	"api.ys1994-user/database"
)

// UserVodCommentReplyAt 影片评论回复@
type UserVodCommentReplyAt struct {
	Base
	ID                    uint                           `json:"id"`
	UserVodCommentReplyID uint                           `json:"-"`
	UserVodCommentReply   userVodCommentReply            `json:"vodCommentReply"`
	Read                  uint8                          `json:"read"`
	UpdatedTime           uint                           `json:"updatedTime"`
	Filter                database.UserVodCommentReplyAt `json:"-"`
}

// UserVodCommentReplyAtList Vod Comment Reply At List
type UserVodCommentReplyAtList struct {
	List
	Data            []UserVodCommentReplyAt        `json:"data"`
	Filter          database.UserVodCommentReplyAt `json:"-"`
}

// Super Super
func (cra *UserVodCommentReplyAt) Super() {
	cra.Base.Super(cra, &cra.Filter)
}

// HaveRead HaveRead
func (cra *UserVodCommentReplyAt) HaveRead(ids []uint) {
	db := database.DB.Model(&database.UserVodCommentReplyAt{}).Where(&cra.Filter)
	if len(ids) > 0 {
		db = db.Where("id IN (?)", ids)
	}
	db.UpdateColumn("read", 2)
}

// Find 列表
func (ral *UserVodCommentReplyAtList) Find() {
	// 不能查删除的
	ral.Filter.DeletedAt = nil
	db := database.DB.Where(&ral.Filter)

	ral.dbHandle(db).Preload("UserVodCommentReply").Find(&ral.Data)

	db.Model(&database.UserVodCommentReplyAt{}).Count(&ral.Count)

	if len(ral.Data) == 0 {
		return
	}

	var users map[uint]*user
	users = make(map[uint]*user)
	var ul userList

	var userVodComments map[uint]*userVodComment
	userVodComments = make(map[uint]*userVodComment)
	var vl userVodCommentList

	for index, cra := range ral.Data {
		if users[cra.UserVodCommentReply.UserID] == nil {
			users[cra.UserVodCommentReply.UserID] = &user{}
		}
		ral.Data[index].UserVodCommentReply.User = users[cra.UserVodCommentReply.UserID]
		ul.Filters.IDs = append(ul.Filters.IDs, cra.UserVodCommentReply.UserID)

		if userVodComments[cra.UserVodCommentReply.UserVodCommentID] == nil {
			userVodComments[cra.UserVodCommentReply.UserVodCommentID] = &userVodComment{}
		}
		ral.Data[index].UserVodCommentReply.UserVodComment = userVodComments[cra.UserVodCommentReply.UserVodCommentID]
		vl.Filters.IDs = append(vl.Filters.IDs, cra.UserVodCommentReply.UserVodCommentID)
	}

	ul.Find()
	for _, user := range ul.Data {
		users[user.ID].ID = user.ID
		users[user.ID].Name = user.Name
	}

	vl.Find()
	for _, userVodComment := range vl.Data {
		userVodComments[userVodComment.ID].ID = userVodComment.ID
		userVodComments[userVodComment.ID].VodID = userVodComment.VodID
		userVodComments[userVodComment.ID].UserID = userVodComment.UserID
	}
}

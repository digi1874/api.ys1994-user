/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-30 15:59:50
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-30 17:15:11
 */

package model

import (
	"api.ys1994-user/database"
)

type userVodCommentReply struct {
	Base
	ID                      uint                    `json:"id"`
	UserVodCommentID       	uint                    `json:"-"`
	UserVodComment          *userVodComment         `json:"vodComment"`
	UserID                  uint                    `json:"-"`
	User                    *user                   `json:"user"`
	Content                 string                  `json:"content"`
}

// UserVodCommentReply 影片评论回复
type UserVodCommentReply struct {
  Base
  ID                           uint                         `json:"id"`
	UserID                       uint                         `json:"userId"`
	Grade                        uint8                        `json:"grade"`
	Content                      string                       `json:"content"`
	Like                         uint                         `json:"like"`
	Dislike                      uint                         `json:"dislike"`
	CreatedTime                  uint                         `json:"createdTime"`
	UserLikeVodCommentReply      UserLikeVodCommentReply      `json:"userLike"`
	Filter                       database.UserVodCommentReply `json:"-"`
}

// UserVodCommentReplyList Vod Comment Reply List
type UserVodCommentReplyList struct {
	List
	Data                          []UserVodCommentReply            `json:"data"`
	Filter                        database.UserVodCommentReply     `json:"-"`
	FilterUserLikeVodCommentReply database.UserLikeVodCommentReply `json:"-"`
}

// Super Super
func (vcr *UserVodCommentReply) Super() {
	vcr.Base.Super(vcr, &vcr.Filter)
}

// IncrementLike Increment Like
func (vcr *UserVodCommentReply) IncrementLike() {
	vcr.Increment("like")
}

// DecrementLike Decrement Like
func (vcr *UserVodCommentReply) DecrementLike() {
	vcr.Decrement("like")
}

// IncrementDislike Increment Dislike
func (vcr *UserVodCommentReply) IncrementDislike() {
	vcr.Increment("dislike")
}

// DecrementDislike Decrement Dislike
func (vcr *UserVodCommentReply) DecrementDislike() {
	vcr.Decrement("dislike")
}

// Find 列表
func (vcl *UserVodCommentReplyList) Find() {
	// 不能查删除的
	vcl.Filter.DeletedAt = nil
	db := database.DB.Where(&vcl.Filter)

	lDB := vcl.dbHandle(db)

	if vcl.FilterUserLikeVodCommentReply.UserID != 0 {
		lDB = lDB.Preload("UserLikeVodCommentReply", "user_id = ?", vcl.FilterUserLikeVodCommentReply.UserID)
	}

	lDB.Find(&vcl.Data)

	db.Model(&database.UserVodCommentReply{}).Count(&vcl.Count)
}

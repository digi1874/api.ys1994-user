/*
 * @Author: lin.zhenhui
 * @Date: 2020-05-10 15:46:35
 * @Last Modified by:   lin.zhenhui
 * @Last Modified time: 2020-05-10 15:46:35
 */

package model

import (
	"api.ys1994-user/database"
)

type userVodCollectionComment struct {
	ID                      uint                    `json:"id"`
	UserID                  uint                    `json:"userId"`
	User                    *user                   `json:"user"`
	UserVodCollectionID     uint                    `json:"vodCollectionID"`
	Content                 string                  `json:"content"`
}

// UserVodCollectionComment 影片集合评论
type UserVodCollectionComment struct {
  Base
  ID                           uint                              `json:"id"`
	UserID                       uint                              `json:"userId"`
	Grade                        uint8                             `json:"grade"`
	Content                      string                            `json:"content"`
	Like                         uint                              `json:"like"`
	Dislike                      uint                              `json:"dislike"`
	CreatedTime                  uint                              `json:"createdTime"`
	UserLikeVodCollectionComment UserLikeVodCollectionComment      `json:"userLike"`
	Filter                       database.UserVodCollectionComment `json:"-"`
}

// UserVodCollectionCommentList 影片集合评论列表
type UserVodCollectionCommentList struct {
	List
	Data                               []UserVodCollectionComment            `json:"data"`
	Filter                             database.UserVodCollectionComment     `json:"-"`
	FilterUserLikeVodCollectionComment database.UserLikeVodCollectionComment `json:"-"`
}

// Super Super
func (vcc *UserVodCollectionComment) Super() {
	vcc.Base.Super(vcc, &vcc.Filter)
}

// IncrementLike Increment Like
func (vcc *UserVodCollectionComment) IncrementLike() {
	vcc.Increment("like")
}

// DecrementLike Decrement Like
func (vcc *UserVodCollectionComment) DecrementLike() {
	vcc.Decrement("like")
}

// IncrementDislike Increment Dislike
func (vcc *UserVodCollectionComment) IncrementDislike() {
	vcc.Increment("dislike")
}

// DecrementDislike Decrement Dislike
func (vcc *UserVodCollectionComment) DecrementDislike() {
	vcc.Decrement("dislike")
}

// Find 列表
func (vcl *UserVodCollectionCommentList) Find() {
	// 不能查删除的
	vcl.Filter.DeletedAt = nil
	db := database.DB.Where(&vcl.Filter)

	lDB := vcl.dbHandle(db)

	if vcl.FilterUserLikeVodCollectionComment.UserID != 0 {
		lDB = lDB.Preload("UserLikeVodCollectionComment", "user_id = ?", vcl.FilterUserLikeVodCollectionComment.UserID)
	}

	lDB.Find(&vcl.Data)

	db.Model(&database.UserVodCollectionComment{}).Count(&vcl.Count)
}

/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-02 14:30:50
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-13 19:40:15
 */

package model

import (
	"api.ys1994-user/database"
)

type userVodComment struct {
	Base
	ID                      uint                    `json:"id"`
	VodID                   uint                    `json:"vodId"`
	Vod                     *vod                    `json:"vod"`
	UserID                  uint                    `json:"userId"`
	User                    *user                   `json:"user"`
	Content                 string                  `json:"content"`
}

type userVodCommentList struct {
	List
	Data         []userVodComment          `json:"data"`
	Filter       database.UserVodComment   `json:"-"`
}

// UserVodComment 影片评论
type UserVodComment struct {
	Base
	ID                      uint                    `json:"id"`
	VodID                   uint                    `json:"vodId"`
	UserID                  uint                    `json:"userId"`
	Grade                   uint8                   `json:"grade"`
	Content                 string                  `json:"content"`
	VodStar                 uint8                   `json:"vodStar"`
	DirectorStar            uint8                   `json:"directorStar"`
	ActorStar               uint8                   `json:"actorStar"`
	ActressStar             uint8                   `json:"actressStar"`
	SuppActorStar           uint8                   `json:"suppActorStar"`
	SuppActressStar         uint8                   `json:"suppActressStar"`
	ScreenplayStar          uint8                   `json:"screenplayStar"`
	CinematographyStar      uint8                   `json:"cinematographyStar"`
	CinematographyStarCount uint                    `json:"cinematographyStarCount"`
	EditStar                uint8                   `json:"editStar"`
	SoundStar               uint8                   `json:"soundStar"`
	VisualStar              uint8                   `json:"visualStar"`
	MakeupStar              uint8                   `json:"makeupStar"`
	CostumeStar             uint8                   `json:"costumeStar"`
	MusicStar               uint8                   `json:"musicStar"`
	Like                    uint                    `json:"like"`
	Dislike                 uint                    `json:"dislike"`
	Reply                   uint                    `json:"reply"`
	CreatedTime             uint                    `json:"createdTime"`
	UserLikeVodComment      UserLikeVodComment      `json:"userLike"`
	Filter                  database.UserVodComment `json:"-"`
}

// UserVodCommentList Vod Comment List
type UserVodCommentList struct {
	List
	Data                     []UserVodComment             `json:"data"`
	Filter                   database.UserVodComment      `json:"-"`
	FilterUserLikeVodComment database.UserLikeVodComment  `json:"-"`
}

// Super Super
func (uvc *UserVodComment) Super() {
	uvc.Base.Super(uvc, &uvc.Filter)
}

// IncrementReply Increment Reply
func (uvc *UserVodComment) IncrementReply() {
	uvc.Increment("reply")
}

// IncrementLike Increment Like
func (uvc *UserVodComment) IncrementLike() {
	uvc.Increment("like")
}

// DecrementLike Decrement Like
func (uvc *UserVodComment) DecrementLike() {
	uvc.Decrement("like")
}

// IncrementDislike Increment Dislike
func (uvc *UserVodComment) IncrementDislike() {
	uvc.Increment("dislike")
}

// DecrementDislike Decrement Dislike
func (uvc *UserVodComment) DecrementDislike() {
	uvc.Decrement("dislike")
}

// Find 列表
func (vcl *UserVodCommentList) Find() {
	// 不能查删除的
	vcl.Filter.DeletedAt = nil
	db := database.DB.Where(&vcl.Filter)

	lDB := vcl.dbHandle(db)

	if vcl.FilterUserLikeVodComment.UserID != 0 {
		lDB = lDB.Preload("UserLikeVodComment", "user_id = ?", vcl.FilterUserLikeVodComment.UserID)
	}

	lDB.Find(&vcl.Data)

	db.Model(&database.UserVodComment{}).Count(&vcl.Count)
}

// Find 列表
func (vcl *userVodCommentList) Find() {
	db := vcl.Filters.dbHandle(database.DB, "user_vod_comments")
	vcl.dbHandle(db.Where(&vcl.Filter)).Find(&vcl.Data)
}

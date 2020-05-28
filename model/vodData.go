/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-02 16:19:22
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 21:09:40
 */

package model

import (
	"api.ys1994-user/database"
)

// VodData 影片互动数据
type VodData struct {
	Base
	VodID                   uint               `json:"vodId"`
	LikeCount               uint               `json:"likeCount"`
	PlayCount               uint               `json:"playCount"`
	VodStar                 uint8              `json:"vodStar"`
	VodStarCount            uint               `json:"vodStarCount"`
	DirectorStar            uint8              `json:"directorStar"`
	DirectorStarCount       uint               `json:"directorStarCount"`
	ActorStar               uint8              `json:"actorStar"`
	ActorStarCount          uint               `json:"actorStarCount"`
	ActressStar             uint8              `json:"actressStar"`
	ActressStarCount        uint               `json:"actressStarCount"`
	SuppActorStar           uint8              `json:"suppActorStar"`
	SuppActorStarCount      uint               `json:"suppActorStarCount"`
	SuppActressStar         uint8              `json:"suppActressStar"`
	SuppActressStarCount    uint               `json:"suppActressStarCount"`
	ScreenplayStar          uint8              `json:"screenplayStar"`
	ScreenplayStarCount     uint               `json:"screenplayStarCount"`
	CinematographyStar      uint8              `json:"cinematographyStar"`
	CinematographyStarCount uint               `json:"cinematographyStarCount"`
	EditStar                uint8              `json:"editStar"`
	EditStarCount           uint               `json:"editStarCount"`
	SoundStar               uint8              `json:"soundStar"`
	SoundStarCount          uint               `json:"soundStarCount"`
	VisualStar              uint8              `json:"visualStar"`
	VisualStarCount         uint               `json:"visualStarCount"`
	MakeupStar              uint8              `json:"makeupStar"`
	MakeupStarCount         uint               `json:"makeupStarCount"`
	CostumeStar             uint8              `json:"costumeStar"`
	CostumeStarCount        uint               `json:"costumeStarCount"`
	MusicStar               uint8              `json:"musicStar"`
	MusicStarCount          uint               `json:"musicStarCount"`
	Filter                  database.VodData   `json:"-"`
}

// UserVodData 一些用户与影片互动数据
type UserVodData struct {
	VodData         VodData            `json:"vodData"`
	UserLikeVod     UserLikeVod        `json:"like"`
	Play            UserPlayM3u8       `json:"play"`
	UserVodComment  UserVodComment     `json:"comment"`
	FilterUser      database.User      `json:"-"`
}

// Super Super
func (vd *VodData) Super() {
	vd.Base.Super(vd, &vd.Filter)
}

// IncrementLike Increment Like
func (vd *VodData) IncrementLike() {
	vd.Increment("like_count")
}

// DecrementLike Decrement Like
func (vd *VodData) DecrementLike() {
	vd.Decrement("like_count")
}

// Detail detail
func (uvd *UserVodData) Detail() {
	uvd.VodData.Super()
	uvd.VodData.Detail()

	if uvd.VodData.VodID == 0 {
		uvd.VodData.Create()
		uvd.VodData.VodID = uvd.VodData.Filter.VodID
	}

	if uvd.FilterUser.ID != 0 {
		uvd.UserLikeVod.Super()
		uvd.UserLikeVod.Filter.UserID = uvd.FilterUser.ID
		uvd.UserLikeVod.Filter.VodID = uvd.VodData.Filter.VodID
		uvd.UserLikeVod.Detail()

		uvd.Play.Super()
		uvd.Play.Filter.UserID = uvd.FilterUser.ID
		uvd.Play.Filter.VodID = uvd.VodData.Filter.VodID
		uvd.Play.Detail()

		uvd.UserVodComment.Super()
		uvd.UserVodComment.Filter.UserID = uvd.FilterUser.ID
		uvd.UserVodComment.Filter.VodID = uvd.VodData.Filter.VodID
		uvd.UserVodComment.Base.Detail()
	}
}

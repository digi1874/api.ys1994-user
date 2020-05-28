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

// VodM3u8Data 影片互动数据
type VodM3u8Data struct {
	Base
	VodM3u8ID       uint                  `json:"vodId"`
	LikeCount       uint                  `json:"likeCount"`
	PlayCount       uint                  `json:"playCount"`
	Filter          database.VodM3u8Data  `json:"-"`
}

// UserVodM3u8Data 一些用户与影片互动数据
type UserVodM3u8Data struct {
	VodM3u8Data     VodM3u8Data        `json:"vodM3u8Data"`
	UserLikeVodM3u8 UserLikeVodM3u8    `json:"like"`
	Play            UserPlayM3u8       `json:"play"`
	// UserVodComment  UserVodComment     `json:"comment"`
	FilterUser      database.User      `json:"-"`
}

// Super Super
func (vd *VodM3u8Data) Super() {
	vd.Base.Super(vd, &vd.Filter)
}

// IncrementLike Increment Like
func (vd *VodM3u8Data) IncrementLike() {
	vd.Increment("like_count")
}

// DecrementLike Decrement Like
func (vd *VodM3u8Data) DecrementLike() {
	vd.Decrement("like_count")
}

// Detail detail
func (uvd *UserVodM3u8Data) Detail() {
	uvd.VodM3u8Data.Super()
	uvd.VodM3u8Data.Detail()

	if uvd.VodM3u8Data.VodM3u8ID == 0 {
		uvd.VodM3u8Data.Create()
		uvd.VodM3u8Data.Detail()
	} else {
		uvd.VodM3u8Data.Increment("play_count")
	}

	if uvd.FilterUser.ID != 0 {
		uvd.UserLikeVodM3u8.Super()
		uvd.UserLikeVodM3u8.Filter.UserID = uvd.FilterUser.ID
		uvd.UserLikeVodM3u8.Filter.VodM3u8ID = uvd.VodM3u8Data.Filter.VodM3u8ID
		uvd.UserLikeVodM3u8.Detail()

		uvd.Play.Super()
		uvd.Play.Filter.UserID = uvd.FilterUser.ID
		uvd.Play.Filter.VodM3u8ID = uvd.VodM3u8Data.Filter.VodM3u8ID
		uvd.Play.Base.Detail()

		// uvd.UserVodComment.Super()
		// uvd.UserVodComment.Filter.UserID = uvd.FilterUser.ID
		// uvd.UserVodComment.Filter.VodM3u8ID = uvd.VodM3u8Data.Filter.VodM3u8ID
		// uvd.UserVodComment.Base.Detail()
	}
}

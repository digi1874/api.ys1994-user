/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-27 19:11:39
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 18:41:33
 */

package model

import (
	"api.ys1994-user/database"
)

// UserLikeVod 收藏视频
type UserLikeVod struct {
	Base
	State           uint8                `json:"state"`
	ID              uint                 `json:"-"`
	Filter          database.UserLikeVod `json:"-"`
}

// UserLikeVodList 收藏视频列表
type UserLikeVodList struct {
	List
	Data            []Vod                 `json:"data"`
	Filter          database.UserLikeVod  `json:"-"`
	FilterVod       database.Vod          `json:"-"`
}

// Super Super
func (ulv *UserLikeVod) Super() {
	ulv.Base.Super(ulv, &ulv.Filter)
}

// Find 列表
func (vll *UserLikeVodList) Find() {
	// 不能查删除的
	vll.FilterVod.DeletedAt = nil
	db := database.DB.Joins("JOIN `user_like_vods` ON `vods`.`id` = `user_like_vods`.`vod_id`").Where(&vll.Filter)
	db = handleFilterVod(db, vll.FilterVod)
	vll.dbHandle(db.Order("`vods`.`updated_time` desc")).Find(&vll.Data)

	db.Model(&database.Vod{}).Count(&vll.Count)
}

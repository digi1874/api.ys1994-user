/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-28 16:04:21
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 21:53:45
 */

package model

import (
	"api.ys1994-user/database"
)

// UserPlayM3u8 视频播放记录
type UserPlayM3u8 struct {
	Base
	ID              uint                  `json:"id"`
	Time            uint                  `json:"time"`
	VodM3u8ID       uint                  `json:"-"`
	VodM3u8         VodM3u8               `json:"m3u8"`
	Filter          database.UserPlayM3u8 `json:"-"`
}

// userPlayM3u8 视频播放记录
type userPlayM3u8 struct {
	UserPlayM3u8
	VodID           uint                 `json:"-"`
	Vod             Vod                  `json:"vod"`
}

// UserPlayM3u8List 视频播放记录列表
type UserPlayM3u8List struct {
	List
	Data            []userPlayM3u8         `json:"data"`
	Filter          database.UserPlayM3u8  `json:"-"`
	FilterVod       database.Vod           `json:"-"`
}

// Super Super
func (upm *UserPlayM3u8) Super() {
	upm.Base.Super(upm, &upm.Filter)
}

// Detail 详情
func (upm *UserPlayM3u8) Detail() {
	database.DB.Where(&upm.Filter).Order("updated_time desc").First(&upm)
	if upm.ID != 0 {
		upm.Related(&upm.VodM3u8)
	}
}

// Find 列表
func (pml *UserPlayM3u8List) Find() {
	// 不能查删除的
	pml.FilterVod.DeletedAt = nil
	db := database.DB.Joins("JOIN `vods` ON `vods`.`id` = `user_play_m3u8`.`vod_id`").Where(&pml.Filter)
	db = handleFilterVod(db, pml.FilterVod)

	pml.dbHandle(db.Order("`user_play_m3u8`.`updated_time` desc")).Preload("Vod").Preload("VodM3u8").Find(&pml.Data)

	db.Model(&database.UserPlayM3u8{}).Count(&pml.Count)
}

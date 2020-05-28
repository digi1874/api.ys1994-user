/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-31 21:31:31
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 22:39:00
 */

package model

import (
	"api.ys1994-user/database"
)

// UserVodCollectionVod 视频集合中的视频
type UserVodCollectionVod struct {
	Base
	ID                      uint                          `json:"id"`
	UserVodCollectionID     uint                          `json:"userVodCollectionID"`
	VodID                   uint                          `json:"vodID"`
	Filter                  database.UserVodCollectionVod `json:"-"`
}

// UserVodCollectionVodList 视频集合中的视频列表
type UserVodCollectionVodList struct {
	List
	Data                    []Vod                          `json:"data"`
	Filter                  database.UserVodCollectionVod  `json:"-"`
	FilterVod               database.Vod                   `json:"-"`
}

// Super Super
func (vcv *UserVodCollectionVod) Super() {
	vcv.Base.Super(vcv, &vcv.Filter)
}

// Delete 删除
func (vcv *UserVodCollectionVod) Delete(vodIDs []uint) {
	database.DB.Where(&vcv.Filter).Where("`vod_id` IN (?)", vodIDs).Delete(&vcv.Filter)
}

// Find 列表
func (cvl *UserVodCollectionVodList) Find() {
	// 不能查删除的
	cvl.FilterVod.DeletedAt = nil
	// cvl.FilterUserVodCollection.DeletedAt = nil
	cvl.Filter.DeletedAt = nil
	db := database.DB.
		Joins("JOIN `user_vod_collection_vods` ON `vods`.`id` = `user_vod_collection_vods`.`vod_id` AND `user_vod_collection_vods`.`deleted_at` IS NULL").
		Where(&cvl.Filter)
  db = handleFilterVod(db, cvl.FilterVod)

  cvl.dbHandle(db.Order("`vods`.`updated_time` desc")).Find(&cvl.Data)

	db.Model(&database.Vod{}).Count(&cvl.Count)
}
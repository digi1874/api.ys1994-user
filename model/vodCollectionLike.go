/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-01 17:42:13
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-01 19:38:21
 */

package model

import (
	"api.ys1994-user/database"
)

// UserLikeVodCollection 收藏视频集合
type UserLikeVodCollection struct {
	Base
	ID              uint
	State           uint8
	Filter          database.UserLikeVodCollection
}

// UserLikeVodCollectionList 收藏视频集合列表
type UserLikeVodCollectionList struct {
	List
	Data                        []UserVodCollection             `json:"data"`
	Filter                      database.UserVodCollection      `json:"-"`
	FilterUserLikeVodCollection database.UserLikeVodCollection  `json:"-"`
}

// Super Super
func (lvc *UserLikeVodCollection) Super() {
	lvc.Base.Super(lvc, &lvc.Filter)
}

// Find 列表
func (vll *UserLikeVodCollectionList) Find() {
	db := database.DB.
		Joins("JOIN `user_like_vod_collections` ON `user_vod_collections`.`id` = `user_like_vod_collections`.`user_vod_collection_id` AND `user_like_vod_collections`.`deleted_at` IS NULL").
		Where(&vll.FilterUserLikeVodCollection)

	db = vll.Filters.dbHandle(db, "user_vod_collections")

	if vll.Filter.Name != "" {
		db = db.Where("`user_vod_collections`.`name` LIKE ?", "%" + vll.Filter.Name + "%")
		vll.Filter.Name = ""
	}
	// 不能查删除的
	vll.Filter.DeletedAt = nil
	db = db.Where(&vll.Filter)

	vll.dbHandle(db.Order("`user_vod_collections`.`updated_time` desc")).Find(&vll.Data)

	db.Model(&database.UserVodCollection{}).Count(&vll.Count)
}

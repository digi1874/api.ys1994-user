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

// UserLikeVodM3u8 收藏视频
type UserLikeVodM3u8 struct {
	Base
	State           uint8                    `json:"state"`
	ID              uint                     `json:"-"`
	Filter          database.UserLikeVodM3u8 `json:"-"`
}

// Super Super
func (ulv *UserLikeVodM3u8) Super() {
	ulv.Base.Super(ulv, &ulv.Filter)
}

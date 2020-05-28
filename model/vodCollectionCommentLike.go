/*
 * @Author: lin.zhenhui
 * @Date: 2020-05-10 15:46:48
 * @Last Modified by:   lin.zhenhui
 * @Last Modified time: 2020-05-10 15:46:48
 */

package model

import (
	"api.ys1994-user/database"
)

// UserLikeVodCollectionComment 影片集合评论点赞贬赞
type UserLikeVodCollectionComment struct {
  Base
  ID                         uint                                  `json:"-"`
	UserVodCollectionCommentID uint                                  `json:"-"`
	State                      uint8                                 `json:"state"`
	Filter                     database.UserLikeVodCollectionComment `json:"-"`
}

// Super Super
func (vcc *UserLikeVodCollectionComment) Super() {
	vcc.Base.Super(vcc, &vcc.Filter)
}

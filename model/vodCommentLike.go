/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-02 14:31:09
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-02 16:05:56
 */

package model

import (
	"api.ys1994-user/database"
)

// UserLikeVodComment 影片评论点赞贬赞
type UserLikeVodComment struct {
  Base
  ID               uint                        `json:"-"`
	UserVodCommentID uint                        `json:"-"`
	State            uint8                       `json:"state"`
	Filter           database.UserLikeVodComment `json:"-"`
}

// Super Super
func (lvc *UserLikeVodComment) Super() {
	lvc.Base.Super(lvc, &lvc.Filter)
}

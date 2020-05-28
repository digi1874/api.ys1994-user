/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-30 16:24:31
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-30 16:25:36
 */

package model

import (
	"api.ys1994-user/database"
)

// UserLikeVodCommentReply 影片评论回复点赞贬赞
type UserLikeVodCommentReply struct {
  Base
  ID                    uint                             `json:"-"`
	UserVodCommentReplyID uint                             `json:"-"`
	State                 uint8                            `json:"state"`
	Filter                database.UserLikeVodCommentReply `json:"-"`
}

// Super Super
func (vcr *UserLikeVodCommentReply) Super() {
	vcr.Base.Super(vcr, &vcr.Filter)
}

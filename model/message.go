/*
 * @Author: lin.zhenhui
 * @Date: 2020-05-13 15:37:33
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-13 15:55:34
 */

package model

import (
	"api.ys1994-user/database"
)

// MessageCount MessageCount
type MessageCount struct {
	VodCommentCount           uint `json:"vodCommentCount"`
	VodCommentReplyCount      uint `json:"vodCommentReplyCount"`
	VodCollectionCommentCount uint `json:"vodCollectionCommentCount"`
	UserID                    uint `json:"-"`
}

// UnreadCount UnreadCount
func (m *MessageCount) UnreadCount() {
	vca := database.UserVodCommentAt{ UserID: m.UserID, Read: 1 }
	database.DB.Model(&database.UserVodCommentAt{}).Where(&vca).Count(&m.VodCommentCount)

	cra := database.UserVodCommentReplyAt{ UserID: m.UserID, Read: 1 }
	database.DB.Model(&database.UserVodCommentReplyAt{}).Where(&cra).Count(&m.VodCommentReplyCount)

	cca := database.UserVodCollectionCommentAt{ UserID: m.UserID, Read: 1 }
	database.DB.Model(&database.UserVodCollectionCommentAt{}).Where(&cca).Count(&m.VodCollectionCommentCount)
}

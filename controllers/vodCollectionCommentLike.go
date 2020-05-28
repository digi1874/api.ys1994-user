/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-27 23:13:38
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-10 16:09:39
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type likeVodCollectionCommentValidator struct {
	VodCollectionCommentID uint  `json:"vodCollectionCommentId" binding:"required,min=1"`
	State                  uint8 `json:"state" binding:"required,min=1,max=3"`
}

// LikeVodCollectionCommentHandle Like Vod CollectionComment
func LikeVodCollectionCommentHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData likeVodCollectionCommentValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	var uvc model.UserVodCollectionComment
	uvc.Super()
	uvc.Filter.ID = formData.VodCollectionCommentID
	uvc.Base.Detail()

	if uvc.ID == 0 {
		cJSONBadRequest(c, "评论不存在")
		return
	}

	var lvc model.UserLikeVodCollectionComment
	lvc.Super()
	lvc.Filter.UserID = userID
	lvc.Filter.UserVodCollectionCommentID = formData.VodCollectionCommentID
	lvc.Detail()

	if lvc.State != formData.State {
		lvc.Filter.State = formData.State
		if lvc.ID == 0 {
			lvc.Create()
		} else {
			lvc.Filter.ID = lvc.ID
			lvc.Update()
		}
		go vodCollectionCommentLikeCount(&uvc, formData.State, lvc.State)
	}

	cJSONOk(c, true)
}

func vodCollectionCommentLikeCount(uvc *model.UserVodCollectionComment, newState, oldState uint8) {
	if newState == oldState { return }

	if oldState == 0 || oldState == 3 {
		if newState == 1 {
			uvc.IncrementLike()
		} else if newState == 2 {
			uvc.IncrementDislike()
		}
	} else if oldState == 1 {
		uvc.DecrementLike()
		if newState == 2 {
			uvc.IncrementDislike()
		}
	} else if oldState == 2 {
		uvc.DecrementDislike()
		if newState == 1 {
			uvc.IncrementLike()
		}
	}
}

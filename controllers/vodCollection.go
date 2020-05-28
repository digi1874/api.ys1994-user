/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-30 22:05:25
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 18:51:53
 */

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
	"api.ys1994-user/utils"
)

type vodCollectionEditValidator struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name" binding:"required,max=50"`
	Pic             string    `json:"pic" binding:"required"`
	Content         string    `json:"content" binding:"required,max=5000"`
	Share           int8      `json:"share" binding:"required,min=1,max=2"`
}

// VodCollectionAddOrEditHandle 新增/编辑集合
func VodCollectionAddOrEditHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	var formData vodCollectionEditValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	var uvc model.UserVodCollection
	uvc.Super()
	if err := utils.Copy(formData, &uvc.Filter); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	uvc.Filter.UserID = userID

	if uvc.Filter.ID == 0 {
		uvc.Create()
		cJSONOk(c, uvc.Filter.ID)
		return
	}

	if errStr, _ := hasVodCollection(uvc.Filter.ID, userID); errStr != "" {
		cJSONBadRequest(c, errStr)
		return
	}

	uvc.Update()
	cJSONOk(c, "修改成功")
}

// DeleteVodCollectionHandle 删除集合
func DeleteVodCollectionHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData deleteValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	var uvc model.UserVodCollection
	uvc.Super()
	uvc.Filter.UserID = userID
	uvc.Delete(formData.IDs)
	cJSONOk(c, "删除成功")
}

// GetVodCollectionListHandle 获取我的视频集合
func GetVodCollectionListHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	if vcl, err := getVodCollectionList(c, userID); err != nil {
		cJSONBadRequest(c, err.Error())
	} else {
		cJSONOk(c, vcl)
	}
}

// GetShareVodCollectionListHandle 获取全部人分享的视频集合
func GetShareVodCollectionListHandle(c *gin.Context) {
	if vcl, err := getVodCollectionList(c, 0); err != nil {
		cJSONBadRequest(c, err.Error())
	} else {
		cJSONOk(c, vcl)
	}
}

// GetVodCollectionDetailHandle 获取我的视频集合详情
func GetVodCollectionDetailHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	if uvc, errStr := getVodCollectionDetail(c, userID); errStr != "" {
		cJSONBadRequest(c, errStr)
	} else {
		cJSONOk(c, uvc)
	}
}

// GetShareVodCollectionDetailHandle 获取分享的视频集合详情
func GetShareVodCollectionDetailHandle(c *gin.Context) {
	if uvc, errStr := getVodCollectionDetail(c, 0); errStr != "" {
		cJSONBadRequest(c, errStr)
	} else {
		cJSONOk(c, uvc)
	}
}

func hasVodCollection(id uint, userID uint) (string, *model.UserVodCollection) {
	var uvc model.UserVodCollection
	uvc.Super()
	uvc.Filter.ID = id
	uvc.Detail()
	if uvc.ID == 0 {
		return "集合不存在", &uvc
	} else if userID != 0 && uvc.UserID != userID {
		return "无权限修改集合", &uvc
	}

	return "", &uvc
}

func getVodCollectionList(c *gin.Context, userID uint) (vcl model.UserVodCollectionList, err error) {
	vcl.Page, vcl.Size, err = listHandle(c, &vcl.Filter, &vcl.Filters, &vcl.Orders)

	if err == nil {
		if userID == 0 {
			vcl.Filter.Share = 1
			vcl.Filter.State = 1
		} else {
			vcl.Filter.UserID = userID
		}
		vcl.Find()
	}

	return vcl, err
}

func getVodCollectionDetail(c *gin.Context, userID uint) (uvc model.UserVodCollection, errStr string) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		return uvc, "id 不能为空"
	}

	uvc.Super()
	if userID == 0 {
		uvc.Filter.Share = 1
		uvc.Filter.State = 1
	} else {
		uvc.Filter.UserID = userID
	}
	uvc.Filter.ID = uint(id)
	uvc.Detail()

	if uvc.ID == 0 {
		return uvc, "集合不存在"
	}

	return uvc, ""
}

func vodCollectionIncrementLike(id uint) {
	var uvc model.UserVodCollection
	uvc.Super()
	uvc.Filter.ID = id
	uvc.IncrementLike()
}

func vodCollectionDecrementLike(id uint) {
	var uvc model.UserVodCollection
	uvc.Super()
	uvc.Filter.ID = id
	uvc.DecrementLike()
}

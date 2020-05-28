/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-31 21:40:58
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 22:42:21
 */

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api.ys1994-user/model"
)

type addVodCollectionVodValidator struct {
	VodIDs            []uint  `json:"vodIds" binding:"required,min=1"`
}

// AddVodCollectionVodHandle 添加集合视频
func AddVodCollectionVodHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	var formData addVodCollectionVodValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	vcID := getVodCollectionID(c, userID)
	if vcID == 0 {
		cJSONBadRequest(c, "集合不存在")
		return
	}

	var ids []uint

	for _, vodID := range formData.VodIDs {
		var vcv model.UserVodCollectionVod
		vcv.Super()
		vcv.Filter.VodID = vodID
		vcv.Filter.UserVodCollectionID = vcID
		vcv.FirstOrCreate()
		ids = append(ids, vcv.Filter.ID)
	}

	cJSONOk(c, ids)
}

// DeleteVodCollectionVodHandle 删除集合视频
func DeleteVodCollectionVodHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	// validator formData
	var formData addVodCollectionVodValidator
	if err := c.ShouldBindJSON(&formData); err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	vcID := getVodCollectionID(c, userID)
	if vcID == 0 {
		cJSONBadRequest(c, "集合不存在")
		return
	}

	var vcv model.UserVodCollectionVod
	vcv.Super()
	vcv.Filter.UserVodCollectionID = vcID
	vcv.Delete(formData.VodIDs)
	cJSONOk(c, "删除成功")
}

// GetVodCollectionVodListHandle 集合视频列表
func GetVodCollectionVodListHandle(c *gin.Context) {
	vcID := getVodCollectionID(c, 0)
	if vcID == 0 {
		cJSONBadRequest(c, "集合不存在")
		return
	}

	var err error
	var cvl model.UserVodCollectionVodList
	cvl.Page, cvl.Size, err = listHandle(c, &cvl.FilterVod, &cvl.Filters, &cvl.Orders)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	cvl.Filter.UserVodCollectionID = uint(vcID)
	cvl.Find()

	cJSONOk(c, cvl)
}

func getVodCollectionID(c *gin.Context, userID uint) (id uint) {
	vodCollectionID, _ := strconv.Atoi(c.Param("id"))
	if vodCollectionID < 1 {
		return 0
	}

	var uvc model.UserVodCollection
	uvc.Super()
	uvc.Filter.UserID = userID
	uvc.Filter.ID = uint(vodCollectionID)
	uvc.Detail()

	return uvc.ID
}
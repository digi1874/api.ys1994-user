/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-02 18:48:37
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-02 18:49:17
 */

package controllers

type deleteValidator struct {
	IDs []uint `json:"ids" binding:"required,min=1"`
}

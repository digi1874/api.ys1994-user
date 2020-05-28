/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-28 16:45:00
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 21:59:01
 */

package model

import (
	"api.ys1994-user/database"
)

// VodM3u8 影片m3u8地址
type VodM3u8 struct {
	Base
	ID              uint             `json:"id"`
	VodID           uint             `json:"-"`
	Name            string           `json:"name"`
	Filter          database.VodM3u8 `json:"-"`
}

// Super Super
func (vm *VodM3u8) Super() {
	vm.Base.Super(vm, &vm.Filter)
}

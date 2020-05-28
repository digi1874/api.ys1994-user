/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-15 16:40:47
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-15 16:48:06
 */

package model

import (
	"api.ys1994-user/database"
)

// UserLogin 登录记录
type UserLogin struct {
	Base
	ID         uint                 `json:"id"`
	UserID     uint                 `json:"-"`
	IP         string               `json:"ip"`
	Exp        uint                 `json:"exp"`
	State      uint8                `json:"state"`
	Filter     database.UserLogin   `json:"-"`
}

// Super Super
func (ul *UserLogin) Super() {
	ul.Base.Super(ul, &ul.Filter)
}

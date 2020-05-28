/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-07 17:23:05
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 18:43:03
 */

package controllers

import (
	"api.ys1994-user/model"
)

func hasVod(id uint) uint {
  if id == 0 {
    return id
  }

  var vod model.Vod
  vod.Super()
  vod.Filter.ID = id
  vod.Detail()
  return vod.ID
}

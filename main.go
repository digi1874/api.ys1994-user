/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-15 13:15:01
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-15 15:50:06
 */

package main

import (
	"api.ys1994-user/database"
	"api.ys1994-user/routers"
)

func main()  {
	defer database.DB.Close()
	routers.Run()
}

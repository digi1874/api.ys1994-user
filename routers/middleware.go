/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-15 16:30:07
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-07 21:47:10
 */

package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var origins = [...]string{"https://ys1994.nl","https://www.ys1994.nl","https://admin.ys1994.nl","https://account.ys1994.nl"}

// func init() {
// 	if process.IsDev {
// 		origins = [...]string{"http://dev.ys1994.nl","http://dev.www.ys1994.nl","http://dev.admin.ys1994.nl","http://dev.account.ys1994.nl"}
// 	}
// }

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		// origin := c.Request.Header.Get("Origin")
		// for _, value := range origins {
		// 	if value == origin {
		// 		c.Header("Access-Control-Allow-Origin", origin)
		// 		break
		// 	}
		// }

		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Auth")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}

/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 15:50:01
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-13 18:07:42
 */

package routers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-user/controllers"
	"api.ys1994-user/process"
)



// Run run Router
func Run() {
	if process.IsDev == false {
		gin.SetMode(gin.ReleaseMode)
	}

	var Router = gin.Default()
	Router.Use(middleware())
	user := Router.Group("/user")

	user.GET("/search", controllers.SearchUserHandle)

	user.GET("/image/*name", controllers.GetImageHandle)
	user.POST("/image", controllers.ImageUploadHandle)

	user.GET("/info", controllers.GetInfoHandle)
	user.POST("/info", controllers.UpdateInfoHandle)
	user.GET("/avatar", controllers.GetAvatarListHandle)

	user.GET("/vodData/:vodId", controllers.GetUserVodDataHandle)
	user.GET("/vodM3u8Data/:vodM3u8Id", controllers.GetUserVodM3u8DataHandle)

	// user.GET("/isLike/vod/:id", controllers.IsLikeVodHandle)   // 预计删除
	user.GET("like/vod", controllers.GetVodLikeListHandle)
	user.POST("like/vod", controllers.LikeVodHandle)
	user.POST("like/vodM3u8", controllers.LikeVodM3u8Handle)

	user.POST("/play/vod", controllers.SaveVodPlayHandle)
	user.DELETE("/play/vod", controllers.DeleteVodPlayHandle)
	user.GET("/play/vod", controllers.GetVodPlayListHandle)
	// user.GET("/play/vod/:vodId", controllers.GetVodPlayDetailHandle)   // 预计删除
	// user.GET("/play/m3u8/:m3u8Id", controllers.GetVodPlayM3u8TimeHandle)   // 预计做 m3u8 data

	user.POST("/vodComment", controllers.SaveVodCommentHandle)
	user.GET("/vodComment", controllers.GetVodCommentListHandle)
	user.GET("/vodComment/:id", controllers.GetVodCommentDetailHandle)
	user.POST("/like/vodComment", controllers.LikeVodCommentHandle)

	user.POST("/vodCommentReply", controllers.SaveVodCommentReplyHandle)
	user.GET("/vodCommentReply", controllers.GetVodCommentReplyListHandle)
	user.POST("/like/vodCommentReply", controllers.LikeVodCommentReplyHandle)

	user.POST("/vodCollection", controllers.VodCollectionAddOrEditHandle)
	user.DELETE("/vodCollection", controllers.DeleteVodCollectionHandle)
	user.GET("/vodCollection", controllers.GetVodCollectionListHandle)
	user.GET("/vodCollection/:id", controllers.GetVodCollectionDetailHandle)
	user.GET("/share/vodCollection", controllers.GetShareVodCollectionListHandle)
	user.GET("/share/vodCollection/:id", controllers.GetShareVodCollectionDetailHandle)

	user.POST("/vodCollection/:id/vod", controllers.AddVodCollectionVodHandle)
	user.DELETE("/vodCollection/:id/vod", controllers.DeleteVodCollectionVodHandle)
	user.GET("/vodCollection/:id/vod", controllers.GetVodCollectionVodListHandle)

	user.GET("/isLike/vodCollection/:id", controllers.IsLikeVodCollectionHandle)
	user.POST("/like/vodCollection", controllers.LikeVodCollectionHandle)
	user.GET("/like/vodCollection", controllers.GetVodCollectionLikeListHandle)

	user.POST("/vodCollectionComment", controllers.SaveVodCollectionCommentHandle)
	user.GET("/vodCollectionComment", controllers.GetVodCollectionCommentListHandle)
	user.POST("/like/vodCollectionComment", controllers.LikeVodCollectionCommentHandle)

	user.GET("/messageUnreadCount", controllers.GetMessageUnreadCount)
	user.GET("/vodCommentAt", controllers.GetVodCommentAtListHandle)
	user.DELETE("/vodCommentAt", controllers.DeleteVodCommentAtHandle)
	user.POST("/read/vodCommentAt", controllers.UpdateVodCommentAtReadHandle)

	user.GET("/vodCommentReplyAt", controllers.GetVodCommentReplyAtListHandle)
	user.DELETE("/vodCommentReplyAt", controllers.DeleteVodCommentReplyAtHandle)
	user.POST("/read/vodCommentReplyAt", controllers.UpdateVodCommentReplyAtReadHandle)

	user.GET("/vodCollectionCommentAt", controllers.GetVodCollectionCommentAtListHandle)
	user.DELETE("/vodCollectionCommentAt", controllers.DeleteVodCollectionCommentAtHandle)
	user.POST("/read/vodCollectionCommentAt", controllers.UpdateVodCollectionCommentAtReadHandle)

	if process.IsDev == false {
		Router.Run("127.0.0.1:8030")
	} else {
		Router.Run("127.0.0.1:8031")
	}
}

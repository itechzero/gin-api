package router

import (
	v1 "gin-api/api/v1"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {
	gin.SetMode(utils.Config.GetString("APP_DEBUG_MODE"))

	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	Router := gin.Default()

	Router.GET("/ping", func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusOK, "pong")
	})

	apiV1 := Router.Group("api/v1")
	{
		initUserRouter(apiV1)
	}
	return Router
}

func initUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("index", v1.UserList)
		UserRouter.POST("register", v1.Register)
		//		api.POST("/v1/users/create", v1.UserCtr.CreateUser)
		//		api.GET("/v1/users/index", v1.UserCtr.GetUserInfo)
		//		api.GET("/v1/users", v1.UserCtr.GetUserList)
		//		api.PUT("/v1/users/update", v1.UserCtr.UpdateEmail)
		//		api.DELETE("/v1/users/delete", v1.UserCtr.DeleteUser)
		//UserRouter.POST("changePassword", v1.ChangePassword)
		//UserRouter.POST("setUserAuthority", v1.SetUserAuthority)
		//UserRouter.DELETE("deleteUser", v1.DeleteUser)
		//UserRouter.PUT("setUserInfo", v1.SetUserInfo)
	}
}

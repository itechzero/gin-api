package v1

import (
	"gin-api/services"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
)

var UserCtr _User

type _User struct {
}

func (user *_User) CreateUser(ctx *gin.Context) {
	data := services.CreateUser(ctx)
	utils.SendResponse(ctx, utils.SUCCESS, data)
}

//func (user *_User) GetUserInfo(ctx *gin.Context) {
//	data := services.GetUserInfo(ctx)
//	utils.SendResponse(ctx, utils.SUCCESS, data)
//}
//
//func (user *_User) GetUserList(ctx *gin.Context) {
//	data := services.GetUserList(ctx)
//	utils.SendResponse(ctx, utils.SUCCESS, data)
//}
//
//func (user *_User) UpdateEmail(ctx *gin.Context) {
//	data := services.UpdateEmail(ctx)
//	utils.SendResponse(ctx, utils.SUCCESS, data)
//}
//
//func (user *_User) DeleteUser(ctx *gin.Context) {
//	data := services.DeleteUser(ctx)
//	utils.SendResponse(ctx, utils.SUCCESS, data)
//}

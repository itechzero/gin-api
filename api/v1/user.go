package v1

import (
	"gin-api/models"
	"gin-api/models/request"
	"gin-api/models/response"
	"gin-api/services"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var R request.Register
	if err := ctx.ShouldBind(&R); err != nil {
		response.FailWithDetailed(nil, err.Error(), ctx)
	}
	user := &models.User{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, userReturn := services.Register(*user)
	if err != nil {
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "success", ctx)
	}
}

func UserList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if err := ctx.ShouldBind(&pageInfo); err != nil {
		response.FailWithDetailed(nil, err.Error(), ctx)
	}
	if err, list, total := services.UserList(pageInfo); err != nil {
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.Paginate{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
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

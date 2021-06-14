package services

import (
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) interface{} {
	//var user models.User
	//if err := ctx.ShouldBind(&user); err != nil {
	//	return err.Error()
	//}
	//db := core.MySQL.GetInstance()
	//if err := db.Create(&user).Error; err != nil {
	//	return err
	//}
	//return user.ID
	return nil
}

//func GetUserInfo(ctx *gin.Context) interface{} {
//	var user models.User
//	if err := ctx.ShouldBind(&user); err != nil {
//		return err.Error()
//	}
//	db := core.ConnectOrm()
//	db.Where(&models.User{Name: user.Name}).First(&user)
//
//	return user
//}
//
//func GetUserList(ctx *gin.Context) interface{} {
//	var user []models.User
//	db := core.ConnectOrm()
//
//	//result := map[string]interface{}{}
//	//db.Model(&models.User{}).Find(&result)
//
//	if err := db.Find(&user).Error; err != nil {
//		return err
//	}
//	return user
//}
//
//func UpdateEmail(ctx *gin.Context) interface{} {
//	var user models.User
//	if err := ctx.ShouldBind(&user); err != nil {
//		return err.Error()
//	}
//	db := core.ConnectOrm()
//	db.Model(&user).Where("name = ?", user.Name).Update("email", user.Email)
//	return user
//}
//
//func DeleteUser(ctx *gin.Context) interface{} {
//	var user models.User
//	if err := ctx.ShouldBind(&user); err != nil {
//		return err.Error()
//	}
//	db := core.ConnectOrm()
//	db.Model(&user).Where("name = ?", user.Name).Delete(&user)
//	return user
//}

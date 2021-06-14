package services

import (
	"errors"
	"gin-api/models"
	"gin-api/models/request"
	"gin-api/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Register(u models.User) (err error, user models.User) {
	if !errors.Is(utils.Gorm.GetInstance().Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), user
	}
	u.Password = utils.Md5V2(u.Password)
	u.UUID = uuid.NewV4()
	err = utils.Gorm.GetInstance().Create(&u).Error
	return
}

func UserList(info request.PageInfo) (err error, list []models.User, total int64) {
	offset := (info.Page - 1) * info.PageSize
	limit := info.PageSize
	db := utils.Gorm.GetInstance().Model(&models.User{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
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

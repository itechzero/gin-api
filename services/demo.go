package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var Demo = &_Demo{}

type _Demo struct {
}

type Form struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (demo *_Demo) GetFormTest(ctx *gin.Context) interface{} {
	var form Form
	if err := ctx.ShouldBind(&form); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return form
}

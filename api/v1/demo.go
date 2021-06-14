package v1

import (
	"gin-api/services"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
)

var DemoCtr _Demo

type _Demo struct {
}

func (demo *_Demo) GetFormTest(ctx *gin.Context) {
	data := services.Demo.GetFormTest(ctx)
	utils.SendResponse(ctx, utils.SUCCESS, data)
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("Origin") != "" {
			switch ctx.Request.Method {
			case http.MethodOptions:
				ctx.Header("Access-Control-Allow-Origin", "*")  // * can be replaced with any domain name which you want support cross-domain
				ctx.Header("Access-Control-Allow-Methods", "*") //支持的跨域请求方法 POST, GET, OPTIONS,DELETE,PUT
				ctx.Header("Access-Control-Allow-Headers", "*") //支持的头信息字段 Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id
			default:
				ctx.Header("Access-Control-Allow-Origin", "*") // * can be replaced with any domain name which you want support cross-domain
			}
		}
		ctx.Next()
	}
}

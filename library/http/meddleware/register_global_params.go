package middleware

import (
	api_http "ihtest/app"
	"ihtest/library/global"

	"git.medlinker.com/golang/xtrace"
	"github.com/gin-gonic/gin"
)

func RegisterGlobalParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 注册自定义上下文
		registerContext(ctx)

		ctx.Next()
	}
}

func registerContext(ctx *gin.Context) {
	// 通过请求Header创建标准的context上下文变量，
	// 注入到已有的MedContext对象中，不影响该对象的原有使用。
	ctx.Set(
		global.HttpCustomContextKey,
		api_http.NewMedContext(xtrace.CreateStdContext(ctx.Request.Header)),
	)
}

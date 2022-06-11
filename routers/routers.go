package routers


//自定义gin框架默认日志中间件

import (
	"GoBlog/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(setting.GinLogger(), setting.GinRecovery(true))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg" : 404,
		})
	})
	return r
}

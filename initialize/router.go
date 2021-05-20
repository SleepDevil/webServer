package initialize

import (
	"yasi_audio/middleware"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	return Router
}

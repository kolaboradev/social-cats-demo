package routers

import (
	"cats_social/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
    router.POST("/v1/user/register", controllers.RegisterUser)
    router.POST("/v1/user/login", controllers.LoginUser)
}
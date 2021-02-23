package routers

import (
	"WAF/api"
	"WAF/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("").Use(middlewares.JWTAuth())
	{
		UserRouter.GET("/info", api.User)
		UserRouter.GET("/check/:id", api.CheckUser)
		UserRouter.POST("/add", api.AddUser)
		UserRouter.POST("/edit", api.EditUser)
		UserRouter.GET("/delete/:id", api.DelUser)
	}
}

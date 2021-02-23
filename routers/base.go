package routers

import (
	"WAF/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	Router.Static("/static", "./static")
	Router.POST("/login", api.Login)
}

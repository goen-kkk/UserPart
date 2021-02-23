package main

import (
	"WAF/middlewares"
	"WAF/models"
	"WAF/routers"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middlewares.Cors())
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ApiGroup := r.Group("/user/v1")
	routers.InitBaseRouter(ApiGroup)
	routers.InitUserRouter(ApiGroup)

	APP := models.Cfg.Section("server")
	if err := r.Run(APP.Key("HOST").String() + ":" + APP.Key("PORT").String()); err != nil {
		log.Fatal(err)
	}
}

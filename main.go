package main

import (
	"app/app/controller"
	"app/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	route := SetupRoutes()

	route.Run(":8080")
}

func SetupRoutes() *gin.Engine {
	route := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := route.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			eg.GET("/hello", controller.HelloWorld)
		}
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return route
}

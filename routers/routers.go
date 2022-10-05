package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"task-sesi8/controllers"
	docs "task-sesi8/docs"
)

func SetupRouter() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	orderController := controllers.NewOrderController()

	docs.SwaggerInfo.Title = "Swagger Assignment 2 Hacktiv8 API"
	docs.SwaggerInfo.Description = "This is a sample server for Assignment 2"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/orders", orderController.CreateOrder)
	r.GET("/orders", orderController.GetOrders)
	r.PUT("/orders/:id", orderController.UpdateOrder)
	r.DELETE("/orders/:id", orderController.DeleteOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//v1 := r.Group("/api/v1")

	r.Run()
}

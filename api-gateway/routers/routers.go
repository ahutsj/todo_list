package routers

import (
	"api-gateway/internal/handler"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouters(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service))

	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// user server
		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)
	}

	return ginRouter

}

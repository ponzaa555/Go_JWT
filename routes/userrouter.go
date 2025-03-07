package routes

import (
	"JWT/middleware"

	"JWT/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {
	route.Use(middleware.Authenticate())
	route.GET("/users", controller.GetUser())
	route.GET("/users/:user_id", controller.GetUser())
}

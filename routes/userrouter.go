package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ponzaa555/Go_JWT/controller"
)

func UserRoute(route *gin.Engine) {
	// route.Use(middleware.Authenticate())
	// route.GET("/users", controller.GetUser())
	route.GET("/users/:user_id", controller.GetUser())
}

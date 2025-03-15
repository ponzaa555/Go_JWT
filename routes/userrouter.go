package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ponzaa555/Go_JWT/controller"
	"github.com/ponzaa555/Go_JWT/middleware"
)

func UserRoute(route *gin.Engine) {
	//route.Use mean will run function middleware.Authenticate() before handle function for path
	route.Use(middleware.Authenticate())
	route.GET("/users", controller.GetUsers())
	route.GET("/users/:user_id", controller.GetUser())
}

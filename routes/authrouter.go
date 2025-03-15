package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ponzaa555/Go_JWT/controller"
)

func AuthRoute(route *gin.Engine) {
	route.POST("/user/signup", controller.SignUp())
	route.POST("/user/login", controller.Login())
}

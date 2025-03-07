package routes

import (
	controller "JWT/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoute(route *gin.Engine) {
	route.POST("/user/signup", controller.SignUp())
	route.POST("/user/login", controller.Login())
}

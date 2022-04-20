package router

import (
	"github.com/cwww3/go-template/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(e *gin.Engine, userController *controller.UserController) {
	g := e.Group("/user")
	g.GET("/:id", userController.GetUser)
	g.POST("", userController.AddUser)
	g.PUT("/:id", userController.ModifyUser)
}

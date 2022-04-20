package controller

import (
	"context"
	"github.com/cwww3/go-template/internal/entity"
	"github.com/cwww3/go-template/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(u usecase.UserUseCase) *UserController {
	return &UserController{
		userUseCase: u,
	}
}

func (uc *UserController) AddUser(c *gin.Context) {
	u, err := uc.userUseCase.AddUser(context.Background(), &entity.User{
		Email: "2239354893@qq.com",
	})
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}
	c.JSON(200, gin.H{"user": u})
}

func (uc *UserController) ModifyUser(c *gin.Context) {
	u, err := uc.userUseCase.ModifyUser(context.Background(), &entity.User{
		ID:    1,
		Email: "2239354893@qq.com",
	})
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}
	c.JSON(200, gin.H{"user": u})
}

func (uc *UserController) GetUser(c *gin.Context) {
	u, err := uc.userUseCase.GetUser(context.Background(), 1)
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}
	c.JSON(200, gin.H{"user": u})
}

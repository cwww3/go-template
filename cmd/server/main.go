package main

import (
	"github.com/cwww3/go-template/internal/controller"
	"github.com/cwww3/go-template/internal/repository/orm"
	"github.com/cwww3/go-template/internal/router"
	"github.com/cwww3/go-template/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	e := gin.Default()
	dsn := ""
	r := orm.NewOrmRepository(dsn)
	uc := usecase.NewUserUseCase(r)
	c := controller.NewUserController(uc)
	router.RegisterUserRoute(e, c)
	log.Fatalln(e.Run(":8080"))
}

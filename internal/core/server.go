package core

import (
	"fmt"
	"github.com/cwww3/go-template/config"
	"github.com/cwww3/go-template/internal/controller"
	"github.com/cwww3/go-template/internal/repository/orm"
	"github.com/cwww3/go-template/internal/router"
	"github.com/cwww3/go-template/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	e *gin.Engine
}

func GetServer() *Server {
	gin.SetMode(config.GetServerConfig().Mode)
	e := gin.Default()
	dsn := config.GetMysqlConfig().GetDsn()
	r := orm.NewOrmRepository(dsn)
	uc := usecase.NewUserUseCase(r)
	c := controller.NewUserController(uc)
	router.RegisterUserRoute(e, c)
	return &Server{e: e}
}

func (s *Server) Start() {
	log.Fatalln(s.e.Run(fmt.Sprintf(":%s", config.GetServerConfig().Port)))
}

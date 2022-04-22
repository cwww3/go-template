package main

import (
	"github.com/cwww3/go-template/config"
	"github.com/cwww3/go-template/internal/core"
	"github.com/cwww3/go-template/pkg/logger"
)

func main() {
	config.LoadConfig("config", "config.dev", "yml")
	config.ParseConfig()
	logger.Init()
	logger.GetLogger().Info("test log")
	s := core.GetServer()
	s.Start()
}

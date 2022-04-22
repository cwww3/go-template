package main

import (
	"github.com/cwww3/go-template/config"
	"github.com/cwww3/go-template/internal/core"
	"github.com/cwww3/go-template/pkg/logger"
	"github.com/spf13/pflag"
)

var configFile string

func main() {
	pflag.StringVarP(&configFile, "config", "c", "config/config.dev.yml", "config path")
	pflag.Parse()
	config.LoadConfig(configFile)
	logger.Init()
	s := core.GetServer()
	s.Start()
}

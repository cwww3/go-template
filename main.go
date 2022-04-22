package main

import (
	"github.com/cwww3/go-template/config"
	"github.com/cwww3/go-template/internal/core"
)

func main() {
	config.LoadConfig("config", "config.dev", "yml")
	config.ParseConfig()
	s := core.GetServer()
	s.Start()
}

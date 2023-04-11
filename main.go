package main

import (
	"github.com/loboothay/gojobs.git/config"
	"github.com/loboothay/gojobs.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	//Initiando o Router
	router.Initialize()
}

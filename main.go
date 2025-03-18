package main

import (
	"github.com/Murilojms7/Login-System/config"
	"github.com/Murilojms7/Login-System/router"
)

var (
	logger config.Logger
)

func main() {
	// Logger
	logger = *config.GetLogger("main")

	// initialize configs
	if err := config.Init(); err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// initialize Router
	router.Initialize()
}

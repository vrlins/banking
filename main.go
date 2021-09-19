package main

import (
	"github.com/vrlins/banking/app"
	"github.com/vrlins/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}

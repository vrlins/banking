package main

import (
	"github.com/vrlins/banking-lib/logger"
	"github.com/vrlins/banking/app"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}

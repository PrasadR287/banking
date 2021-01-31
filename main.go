package main

import (
	"github.com/PrasadR287/banking/app"
	"github.com/PrasadR287/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}

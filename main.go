package main

import (
	"github.com/sanran4/banking/app"
	"github.com/sanran4/banking/logger"
)

func main() {
	logger.Info("Starting application")
	app.Start()
}

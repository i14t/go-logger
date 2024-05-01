package main

import (
	"github.com/i14t/go-logger"
)

func main() {
	opts := logger.LoggerOptions{
		Level:      "debug",
		Name:       "DEMO",
		HideCaller: false,
	}
	log := logger.NewLogger(opts)
	log.Debug("I am a debug log")
	log.Info("I am a info log")
	log.Warn("I am a warn log")
	log.Error("I am a error log")
}

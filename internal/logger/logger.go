package logger

import (
	"log"
	"os"
)

func InitializeLogger() *log.Logger {

	file, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := log.New(file, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.SetOutput(file)

	return logger
}

package logger

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	return logger
}

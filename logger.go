package logger

import (
	"log"
	"os"
)

// CreateLogger creates a logger that writes to the given filename
func CreateLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

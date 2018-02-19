package logger

import (
	"fmt"
	"log"
	"os"
)

// CreateLogger creates a logger that writes to the given filename
func CreateLogger(filename string) (*log.Logger, error) {
	file, err := os.OpenFile(filename+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		err = fmt.Errorf("error in creating logger file: %x", err)
		return nil, err
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	return logger, nil
}

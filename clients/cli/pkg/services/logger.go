package services

import (
	"encoding/json"

	"fmt"

	"github.com/golang/glog"
)

// Logger logger.
type Logger struct {
}

// Info prints the provided data.
func (logger *Logger) Info(data interface{}) {
	formatedData, err := logger.getFormatedJSON(data)
	if err != nil {
		fmt.Println(data)
	} else {
		fmt.Println(formatedData)
	}
}

// Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func (logger *Logger) Exit(args ...interface{}) {
	glog.Exit(args)
}

func (logger *Logger) getFormatedJSON(data interface{}) (string, error) {
	formattedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(formattedData), nil
}

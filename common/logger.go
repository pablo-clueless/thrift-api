package common

import (
	"fmt"
)

type LogLevel string

const (
	Info  LogLevel = "INFO"
	Warn  LogLevel = "WARN"
	Error LogLevel = "ERROR"
)

func Logger(level LogLevel, message string, err error) {
	logMessage := fmt.Sprintf("[%s] %s", level, message)

	if err != nil {
		logMessage = fmt.Sprintf("%s: %v", logMessage, err)
	}

	fmt.Println(logMessage)
}

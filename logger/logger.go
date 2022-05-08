package logger

import (
	"log"
)

const (
	DebugLevel = LogLevel(iota)
	InfoLevel
	WarningLevel
	ErrorLevel
)

var logLevel = InfoLevel

type LogLevel int

//初始化
func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func SetLevel(level LogLevel) {
	logLevel = level
}

type Logger interface {
	PrintLog(format string, prefix string, v ...interface{})
}

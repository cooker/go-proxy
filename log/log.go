package log

import (
	"fmt"
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

func SetLogLevel(level LogLevel) {
	logLevel = level
}

func writeLog(level LogLevel, prefix, format string, v ...interface{}) string {
	if level < logLevel {
		return ""
	}
	var data string
	var err error
	if v == nil || len(v) == 0 {
		data = format
	} else {
		data = fmt.Sprintf(format, v...)
		er, isError := v[len(v)-1].(error)
		if isError {
			err = er
		}
	}
	log.Print(prefix + data)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return data
}

func Debug(format string, v ...interface{}) {
	writeLog(DebugLevel, "[Debug]", format, v...)
}

func Info(format string, v ...interface{}) {
	writeLog(InfoLevel, "[Info]", format, v...)
}

func Warning(format string, v ...interface{}) {
	writeLog(WarningLevel, "[Warning]", format, v...)
}

func Error(format string, v ...interface{}) {
	writeLog(ErrorLevel, "[Error]", format, v...)
}

package logger

import (
	"fmt"
	"log"
	"strings"
)

type DefaultLogger struct {
	Logger
}

func (*DefaultLogger) PrintLog(format string, prefix string, v ...interface{}) {
	if v == nil || len(v) == 0 {
		log.Println(prefix, format)
	} else {
		verr := v[len(v)-1]
		e, ok := verr.(error)
		if ok {
			var vs []interface{}
			vs = append(vs, v[:len(v)-1]...)
			vs = append(vs, "")
			if strings.Contains(format, "%") {
				log.Println(prefix, fmt.Sprintf(format, vs...), "异常：", e)
			} else {
				log.Println(prefix, format, "异常：", e)
			}
		} else {
			log.Println(prefix, fmt.Sprintf(format, v...))
		}
	}
}

func (l *DefaultLogger) Debug(format string, v ...interface{}) {
	if DebugLevel < logLevel {
		return
	}
	l.PrintLog(format, "[Debug]", v...)
}

func (l *DefaultLogger) Info(format string, v ...interface{}) {
	if InfoLevel < logLevel {
		return
	}
	l.PrintLog(format, "[Info]", v...)
}

func (l *DefaultLogger) Warn(format string, v ...interface{}) {
	if WarningLevel < logLevel {
		return
	}
	l.PrintLog(format, "[Warn]", v...)
}

func (l *DefaultLogger) Error(format string, v ...interface{}) {
	if ErrorLevel < logLevel {
		return
	}
	l.PrintLog(format, "[Error]", v...)
}

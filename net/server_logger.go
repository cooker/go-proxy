package net

import (
	"cooker/go-proxy/logger"
	"fmt"
	"strings"
)

var LOG *ServerLogger

type ServerLogger struct {
	logger.DefaultLogger
	Server *Server
}

func (l *ServerLogger) Debug(format string, v ...interface{}) {
	l.DefaultLogger.Debug(serverReqId(l.Server.GetReqId(), format), v...)
}

func (l *ServerLogger) Info(format string, v ...interface{}) {
	l.DefaultLogger.Info(serverReqId(l.Server.GetReqId(), format), v...)
}

func (l *ServerLogger) Warn(format string, v ...interface{}) {
	l.DefaultLogger.Warn(serverReqId(l.Server.GetReqId(), format), v...)
}

func (l *ServerLogger) Error(format string, v ...interface{}) {
	l.DefaultLogger.Error(serverReqId(l.Server.GetReqId(), format), v...)
}

func serverReqId(reqId int64, format string) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("[%.3d] ", reqId))
	builder.WriteString(format)
	return builder.String()
}

//初始化空logger
func init() {
	LOG = new(ServerLogger)
	LOG.DefaultLogger = logger.DefaultLogger{}
	LOG.Server = &Server{}
}

package testing

import (
	"cooker/go-proxy/logger"
	"cooker/go-proxy/net"
	"errors"
	"fmt"
	log2 "log"
	"runtime"
	"testing"
)

func TestLogInfo(t *testing.T) {
	net.LOG.Info("测试")
	net.LOG.Info("测试 %s", errors.New("sasa"))
}

func TestLog(t *testing.T) {
	log2.SetFlags(log2.Ldate | log2.Ltime | log2.Lshortfile)
	log2.Println("[INFO]", "sasa", "csasa")
	err := errors.New("sasa")
	err = fmt.Errorf("%v", err)
	stack := make([]uintptr, 50)
	_ = runtime.Callers(1, stack[:])
	frames := runtime.CallersFrames(stack)
	for next, more := frames.Next(); more; {
		println(next.Function, next.File, next.Line)
	}
}

func TestLogInit(t *testing.T) {
	logger.SetLevel(logger.InfoLevel)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log2.Println("sa", "sa")
	logg := new(logger.DefaultLogger)
	logg.Info("sasa %s %s", "sas", errors.New("sasa"))
	logg.Error("sasa %s %s", "sas", errors.New("sasa"))
}

func TestServerLog(t *testing.T) {
	net.LOG = new(net.ServerLogger)
	net.LOG.DefaultLogger = logger.DefaultLogger{}
	net.LOG.Server = &net.Server{}

	net.LOG.Info("sasa")
}

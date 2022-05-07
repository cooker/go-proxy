package testing

import (
	"cooker/go-proxy/log"
	"errors"
	"testing"
)

func TestLogInfo(t *testing.T) {
	log.Info("测试")
	log.Info("测试 %s", errors.New("sasa"))
}

func TestLog(t *testing.T) {

}

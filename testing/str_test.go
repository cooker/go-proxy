package testing

import (
	"cooker/go-proxy/core"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

func TestScanf(t *testing.T) {
	var str, abc string
	fmt.Sscanf("sa 123", "%s%s", &str, &abc)
	println(str)
	println(abc)
}

func TestByteLen(t *testing.T) {
	var b []byte = []byte("hello")
	println(len(b))
}

func Test2Json(t *testing.T) {
	config := core.NewConfig(123, 6)
	json, err := json.Marshal(config)
	if err != nil {
		println(err)
	}
	println(string(json))
}

func TestByte2Str(t *testing.T) {

	builder := new(strings.Builder)
	builder.Write([]byte("sasa"))
	println(unsafe.Pointer(&builder))
	builder.Write([]byte("哇撒"))
	println(unsafe.Pointer(&builder))
	println("方案一：", builder.String())

	b := []byte("sasa")

	pointer := *(*string)(unsafe.Pointer(&b))
	println(pointer)
}

package main

import (
	"cooker/go-proxy/core"
	"cooker/go-proxy/log"
	"cooker/go-proxy/utils"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestConfigWrite(t *testing.T) {
	var cfg = core.NewConfig(13800, 6)
	data, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		log.Error("Json 序列化失败 %s", err)
	}
	println(string(data))
	ioutil.WriteFile("config.json", data, 0644)
}

func TestIsExists(t *testing.T) {
	_, err := os.Stat(utils.GEO_FILE)
	println(err == nil) // 返回nil，文件存在
}

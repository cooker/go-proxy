package main

import (
	"cooker/go-proxy/core"
	"cooker/go-proxy/log"
	"cooker/go-proxy/net"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	version = flag.Bool("version", false, "显示版本")
	cfg     = flag.String("c", "config.json", "配置文件")
)

var logLevels map[string]log.LogLevel

func main() {
	initLogLevel()
	flag.Parse()
	if *version {
		fmt.Printf("App info %s (%s) %s", core.AppName, core.Version, core.Desc)
		fmt.Println()
		return
	}
	file, err := os.Open(*cfg)
	if err != nil {
		log.Error("配置文件，读取出错 %s", err)
		return
	}
	var config core.Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Error("配置文件，反序列化出错 %s", err)
		return
	}

	server := net.GetServer(&config)
	server.Start()
}

func initLogLevel() {
	logLevels = make(map[string]log.LogLevel)

	logLevels["debug"] = log.DebugLevel
	logLevels["info"] = log.InfoLevel
	logLevels["warn"] = log.WarningLevel
	logLevels["error"] = log.ErrorLevel
}

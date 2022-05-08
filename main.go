package main

import (
	"cooker/go-proxy/core"
	"cooker/go-proxy/logger"
	"cooker/go-proxy/net"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	version = flag.Bool("version", false, "显示版本")
	cfg     = flag.String("c", "config.json", "配置文件")
)

var logLevels map[string]logger.LogLevel

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
		log.Printf("配置文件，读取出错 %s", err)
		return
	}
	var config core.Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Printf("配置文件，解析出错 %s", err)
		return
	}
	logger.SetLevel(logLevels[config.Loglevel])
	server := net.GetServer(&config)
	server.Start()
}

func initLogLevel() {
	logLevels = make(map[string]logger.LogLevel)

	logLevels["debug"] = logger.DebugLevel
	logLevels["info"] = logger.InfoLevel
	logLevels["warn"] = logger.WarningLevel
	logLevels["error"] = logger.ErrorLevel
}

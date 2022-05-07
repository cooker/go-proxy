package net

import (
	"cooker/go-proxy/core"
	"cooker/go-proxy/log"
	"cooker/go-proxy/utils"
	"fmt"
	"net"
	"time"
)

type Server struct {
	config *core.Config
	reqId  int64
}

func (server Server) getPort() int {
	return server.config.Port
}

func (server Server) setPort(port int) {
	server.config.Port = port
}

func (server Server) GetConnectTimeOut() time.Duration {
	return server.config.ConnTimeOut
}

func GetServer(config *core.Config) Server {
	server := Server{config: config}
	return server
}

func (server *Server) Start() {
	err := server.check()
	if err != nil {
		return
	}
	log.Info("启动服务，端口: %d", server.getPort())
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", server.getPort()))
	if err != nil {
		log.Error("启动服务，出错", err)
		return
	}
	for {
		client, err := listen.Accept()
		server.reqId++

		if err != nil {
			log.Error("[%.3d] 接收请求，出错 %s", server.reqId, err)
		}
		log.Info("[%.3d] 接收请求：%s 定位：%s", server.reqId, client.RemoteAddr(), utils.GpsIp(client.RemoteAddr().(*net.TCPAddr).IP.String()))
		proxy := newProxyConn(server.reqId, client, server)
		go func() {
			err := proxy.handle()
			if err != nil {
				log.Error("[%.3d] 代理处理失败", server.reqId, err)
				proxy.Close()
			}
		}()
	}
}

func (server *Server) check() error {
	if server.getPort() == 0 {
		port, err := utils.GetFreePort()
		if err != nil {
			log.Error("服务启动异常", err)
			return err
		}

		server.setPort(port)
	}
	return nil
}

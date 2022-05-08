package net

import (
	"cooker/go-proxy/core"
	"cooker/go-proxy/logger"
	"cooker/go-proxy/utils"
	"fmt"
	"net"
	"time"
)

type Server struct {
	config *core.Config
	reqId  int64
}

func (server Server) GetReqId() int64 {
	return server.reqId
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
	LOG = new(ServerLogger)
	LOG.DefaultLogger = logger.DefaultLogger{}
	LOG.Server = server

	err := server.check()
	if err != nil {
		return
	}
	LOG.Info("启动服务，端口: %d", server.getPort())
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", server.getPort()))
	if err != nil {
		LOG.Error("启动服务，出错", err)
		return
	}
	for {
		client, err := listen.Accept()
		server.reqId++

		if err != nil {
			LOG.Error("接收请求，出错 %s", err)
		}
		LOG.Info("client：%s 定位：%s", client.RemoteAddr(), utils.GpsIp(client.RemoteAddr().(*net.TCPAddr).IP.String()))
		proxy := newProxyConn(server.reqId, client, server)
		go func() {
			err := proxy.handle()
			if err != nil {
				LOG.Error("代理处理失败", err)
				proxy.Close()
			}
		}()
	}
}

func (server *Server) check() error {
	if server.getPort() == 0 {
		port, err := utils.GetFreePort()
		if err != nil {
			LOG.Error("服务启动异常", err)
			return err
		}

		server.setPort(port)
	}
	return nil
}

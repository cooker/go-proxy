package net

import (
	"bufio"
	"cooker/go-proxy/core"
	"cooker/go-proxy/utils"
	"fmt"
	"io"
	net2 "net"
	"net/http"
	"sync"
	"time"
)

type ProxyConn struct {
	net2.Conn
	proxy *Server
	reqId int64
}

func newProxyConn(reqId int64, conn net2.Conn, proxy *Server) *ProxyConn {
	return &ProxyConn{
		Conn:  conn,
		reqId: reqId,
		proxy: proxy,
	}
}

func (p *ProxyConn) handle() error {
	reader := core.NewHistoryReader(p.Conn)
	clientReader := bufio.NewReader(reader)
	request, err := http.ReadRequest(clientReader)
	if err != nil {
		return fmt.Errorf("http.ReadRequest %w", err)
	}

	//logger.Info("消息请求：\n%s", string(reader.HistoryBuffer()))
	hostPort := utils.HostPort(request)
	LOG.Info("请求地址：%s", hostPort)
	if request.Method == "CONNECT" {
		LOG.Info("CONNECT %s", hostPort)
		p.keepalive(hostPort)
		return nil
	}
	LOG.Info("request %v", request.URL.Path)
	p.noKeepAlive(hostPort, reader.HistoryBuffer())
	return nil
}

func (p *ProxyConn) keepalive(hostport string) {
	clientConn := p.Conn
	targetConn, err := net2.DialTimeout("tcp", hostport, p.proxy.GetConnectTimeOut()*time.Second)
	if err != nil {
		core.SimpleResponse(http.StatusBadGateway).Write(clientConn)

		if err := clientConn.Close(); err != nil {
			LOG.Error("关闭连接失败", err)
		}
		return
	}
	core.SimpleResponse(http.StatusOK).Write(clientConn)
	//io.WriteString(clientConn, core.CONTENT_OK)
	copyDataAndClose(clientConn, targetConn)
}

func (p *ProxyConn) noKeepAlive(hostPort string, reqData []byte) {
	clientConn := p.Conn

	targetConn, err := net2.DialTimeout("tcp", hostPort, p.proxy.GetConnectTimeOut()*time.Second)
	if err != nil {
		LOG.Error("转发请求，失败 %v", err)
		return
	}
	targetConn.Write(reqData)
	copyDataAndClose(clientConn, targetConn)
}

func copyDataAndClose(a, b net2.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)
	go copyData(a, b, &wg)
	go copyData(b, a, &wg)
	wg.Wait()
	a.Close()
	b.Close()
}

func copyData(a net2.Conn, b net2.Conn, wg *sync.WaitGroup) {
	if _, err := io.Copy(b, a); err != nil {
		LOG.Warn("数据拷贝，出错", err)
	}
	wg.Done()
}

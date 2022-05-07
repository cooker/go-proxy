package net

import (
	"bufio"
	"cooker/go-proxy/core"
	"cooker/go-proxy/log"
	"cooker/go-proxy/utils"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

type ProxyConn struct {
	net.Conn
	proxy *Server

	reqId int64
}

func newProxyConn(reqId int64, conn net.Conn, proxy *Server) *ProxyConn {
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

	//log.Info("消息请求：\n%s", string(reader.HistoryBuffer()))
	hostPort := utils.HostPort(request)
	log.Info("[%.3d] 请求地址：%s", p.reqId, hostPort)
	if request.Method == "CONNECT" {
		log.Info("[%.3d] CONNECT %s", p.reqId, hostPort)
		p.keepalive(hostPort)
		return nil
	}
	log.Info("[%.3d] request %v", request.URL.Path)
	p.noKeepAlive(hostPort, reader.HistoryBuffer())
	return nil
}

func (p *ProxyConn) keepalive(hostport string) {
	clientConn := p.Conn
	targetConn, err := net.DialTimeout("tcp", hostport, p.proxy.GetConnectTimeOut()*time.Second)
	if err != nil {
		s := "HTTP/1.1 502 Bad Gateway\r\n\r\n"
		if _, err := io.WriteString(clientConn, s); err != nil {
			log.Error("[%.3d]转发失败 %s", p.reqId, err)
		}
		if err := clientConn.Close(); err != nil {
			log.Error("[%.3d]关闭连接失败 %s", p.reqId, err)
		}
		return
	}
	clientConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

	copyDataAndClose(clientConn, targetConn)
}

func (p *ProxyConn) noKeepAlive(hostPort string, reqData []byte) {
	clientConn := p.Conn

	targetConn, err := net.DialTimeout("tcp", hostPort, p.proxy.GetConnectTimeOut()*time.Second)
	if err != nil {
		log.Error("转发请求，失败 %v", err)
		return
	}
	targetConn.Write(reqData)
	copyDataAndClose(targetConn, clientConn)
}

func copyDataAndClose(a, b net.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)
	go copyData(a, b, &wg)
	go copyData(b, a, &wg)
	wg.Wait()
	a.Close()
	b.Close()
}

func copyData(a net.Conn, b net.Conn, wg *sync.WaitGroup) {
	if _, err := io.Copy(b, a); err != nil {
		log.Warning("数据转发，出错", err)
	}
	wg.Done()
}

package utils

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"unsafe"
)

var hasPort = regexp.MustCompile(`:\d+$`)
var portMap = map[string]string{
	"http":  "80",
	"https": "443",
}

func GetFreePort() (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port, nil
}

func HostPort(r *http.Request) (host string) {
	host = r.URL.Host
	if !hasPort.MatchString(host) {
		host = net.JoinHostPort(host, portMap[r.URL.Scheme])
	}
	return host
}

const GEO_FILE = "GeoLite2-City.mmdb"

func GpsIp(ip string) string {
	_, err := os.Stat(GEO_FILE)
	if err == nil {
		return ""
	}
	db, err := geoip2.Open(GEO_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	city, err := db.City(net.ParseIP(ip))
	if err != nil || city == nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s %s", city.Country.Names["zh-CN"], city.City.Names["zh-CN"])
}

func ToString(datas *[]byte) string {
	return *(*string)(unsafe.Pointer(datas))
}

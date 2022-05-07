package core

import "time"

type Config struct {
	Port        int           `json:"port"`
	Loglevel    string        `json:"loglevel"`
	ConnTimeOut time.Duration `json:"connTimeOut"`
}

func NewConfig(port int, conntimeout time.Duration) Config {
	return Config{Port: port, Loglevel: "Info", ConnTimeOut: conntimeout}
}

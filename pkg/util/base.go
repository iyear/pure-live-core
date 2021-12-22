package util

import (
	"fmt"
	"golang.org/x/net/proxy"
	"net"
	"strings"
)

func IF(f bool, a interface{}, b interface{}) interface{} {
	if f {
		return a
	}
	return b
}

func GetBetweenString(s, start, end string) string {
	if len(s) == 0 {
		return ""
	}
	if start == "" {
		return s
	}
	if end == "" {
		return s
	}
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return ""
	}
	startIndex += len(start)
	endIndex := strings.Index(s, end)
	if endIndex == -1 {
		return ""
	}
	return s[startIndex:endIndex]
}

func MustGetSocks5(host string, port int, user, password string) proxy.Dialer {
	if host == "" || port == 0 {
		return &net.Dialer{}
	}
	dialer, err := proxy.SOCKS5("tcp", fmt.Sprintf("%s:%d", host, port), &proxy.Auth{
		User:     user,
		Password: password,
	}, proxy.Direct)
	if err != nil {
		return &net.Dialer{}
	}
	return dialer
}

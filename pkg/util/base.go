package util

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/net/proxy"
	"math/rand"
	"net"
	"strings"
)

func IF(f bool, a interface{}, b interface{}) interface{} {
	if f {
		return a
	}
	return b
}

// GetBetweenString returns the string between the start and end strings.
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

// MustGetSocks5 returns a socks5 proxy.
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

const letters = "1234567890abcdefghijklmnopqrstuvwxyz"

func RandLetters(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func BigEndianUint16(v uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, v)
	return buf
}

func BigEndianUint32(v uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, v)
	return buf
}

func BigEndianUint64(v uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, v)
	return buf
}

func PutBytes(bytes ...[]byte) []byte {
	var buf []byte
	for _, b := range bytes {
		buf = append(buf, b...)
	}
	return buf
}

func GetCookie(cookies string, key string) string {
	if cookies == "" {
		return ""
	}
	if key == "" {
		return ""
	}
	cookieArr := strings.Split(cookies, ";")
	for _, cookie := range cookieArr {
		t := strings.Split(strings.TrimSpace(cookie), "=")
		if t[0] == key {
			return t[1]
		}
	}
	return ""
}

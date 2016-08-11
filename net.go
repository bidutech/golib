package SHcommon

import (
	"net"
)

func IsIp(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return false
	}
	return true
}

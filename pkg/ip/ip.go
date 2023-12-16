package ip

import (
	"net"
)

func GetLocalIP() string {
	conn, err := net.Dial("udp", "www.baidu.com:80")
	if err != nil {
		return ""
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

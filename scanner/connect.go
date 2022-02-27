package scanner

import (
	"fmt"
	"net"
	"time"

	"PortScan/vars"
)

func Connect(ip string, port int) (string, int, error) {

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), time.Duration(vars.Timeout)*time.Second)

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	return ip, port, err
}

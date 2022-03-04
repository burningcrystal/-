package scanners

import (
	"fmt"
	"net"
	"time"
)

func TcpConnect(ip net.IP,port int) (net.Conn,error) {
	conn,err := net.DialTimeout("tcp",fmt.Sprintf("%v:%v",ip,port),1*time.Second)
	defer func(){
		if conn != nil {
			_ = conn.Close()
		}
	}()
	return conn,err
}

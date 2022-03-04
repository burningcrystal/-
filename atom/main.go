package main

import (
	"atom/scanners"
	"atom/utils"
	"fmt"
	"os"
)

func start_scan(){
	if len(os.Args) == 3 {
		ipList := os.Args[1]
		portList := os.Args[2]
		ips,err := utils.GetIpList(ipList)
		_ = err
		ports,err := utils.GetPorts(portList)
		for _,ip := range ips {
			for _,port := range ports {
				_,err := scanners.TcpConnect(ip,port)
				if err != nil {
					continue
				}
				fmt.Printf("检测到ip:'%s',port:'%v'开放\n",ip,port)
			}
		}
	} else {
		fmt.Printf("%v参数错误,需要依次传入ipList与ports\n",os.Args[0])
	}
}

func main() {
	start_scan()
}


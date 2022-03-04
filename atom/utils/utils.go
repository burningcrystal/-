package utils

import (
	"atom/vars"
	"fmt"
	"github.com/malfunkt/iprange"
	"net"
	"strconv"
	"strings"
)
/*  先填个坑，回头整体功能做完在回来优化
type rootNode struct {
	childNode [11]*node
	end bool
}
type node struct {
	childNode [11]*node
	element string
	end bool
}

func BuildIpTree(ipList []net.IP){
	root := new(rootNode)
	root.end = false
	for i,v := range ipList{
		for j,k := range v {

		}
	}
}*/
//GetFileAsStr用来读取以逗号为分隔符的文件内容，生成一个字符串切片并返回。
func GetIpList(iplist string) ([]net.IP,error) {
	addressRange,err := iprange.ParseList(iplist)
	if err != nil {
		return nil,err
	}
	ips := addressRange.Expand()
	return ips,err
}

func GetPorts(portlist string)([]int,error) {
	ports := []int{}
	if portlist == "" {
		return ports,nil
	}
	ranges := strings.Split(portlist,",")
	for _,r := range ranges {
		r = strings.TrimSpace(r)  //去掉每个ip首尾的空格
		if strings.Contains(r,"-"){
			parts := strings.Split(r,"-")
			if len(parts)!=2 {
				return nil,fmt.Errorf("端口段输入非法！要求输入两个数字，用-连接.'%s'",r)
			}

			p1,err := strconv.Atoi(parts[0])   //将开头的端口字符转换成数字
			if err!= nil {
				return nil, fmt.Errorf("非法的起始端口数字'%s'", p1)
			}
			p2,err := strconv.Atoi(parts[1])   //将开头的端口字符转换成数字
			if err!= nil{
				return nil,fmt.Errorf("非法的结束端口数字'%s'",p2)
			}
			if p2<p1 {
				return nil,fmt.Errorf("端口段逻辑错误！'%s-%s',或许是'%s-%s'",p1,p2,p2,p1)
			}
			for i:=p1;i<=p2;i++{
				ports=append(ports,i)
			}
		} else {
			if port,err := strconv.Atoi(r);err!=nil{
				return nil,fmt.Errorf("端口输入非法！'%s'",port)
			} else {
				ports = append(ports,port)
			}
		}
	}
	return ports,nil
}

func GetTaskList(ipList []net.IP,ports []int)([]map[string]int,int){   //返回ip与端口的映射表和表的长度
	tasks := make([]map[string]int,0)
	for _,ip := range ipList {
		for _,port := range ports {
			tmp := map[string]int{ip.String():port}
			tasks = append(tasks,tmp)
		}
	}

	return tasks,len(tasks)
}

func AssignTask(tasks []map[string]int) {
	batchs := len(tasks)/vars.ThreadNumber
	for i=0;i<batchs;i++ {

	}
}

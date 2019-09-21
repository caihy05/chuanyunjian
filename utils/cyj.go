package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var listenPort = flag.String("listen", "", "usage: -listen port  example: cyj -listen 8888")
var tranPort = flag.String("tran", "", "usage: -tran ip1:port1 ip2:port2  example: cyj -tran 127.0.0.1:8888 192.168.0.100:30000")

func init() {
	flag.Parse()
	err := CheckPort(*listenPort)
	if err != nil {
		os.Exit(1)
	}
	ipPorts := strings.Split(*tranPort, "")
	for _, ipPort := range ipPorts {
		err = CheckIp(strings.Split(ipPort, ":")[0])
		if err != nil {
			os.Exit(1)
		}
		err = CheckPort(strings.Split(ipPort, ":")[1])
		if err != nil {
			os.Exit(1)
		}
	}
	fmt.Println("------------------------it is ok!----------------------")
}
func Cyj() {

}

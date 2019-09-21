package utils

import (
	"flag"
	"fmt"
)

var listenPort = flag.String("listen","","usage: -listen port  example: cyj -listen 8888")
var tranPort = flag.String("tran","","usage: -tran ip1:port1 ip2:port2  example: cyj -tran 127.0.0.1:8888 192.168.0.100:30000")

func init()  {

	PrintWelcome()
	fmt.Println("------------------------it is ok!----------------------")
}
func Cyj(){


}



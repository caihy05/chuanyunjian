package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var listenPort = flag.String("listen", "", "usage: -listen port  example: cyj -listen 28888")
var monitorPort = flag.String("monitor", "", "usage: -monitor port1,port2  example: cyj -monitor 28888,38888")
var tranPort = flag.String("tran", "", "usage: -tran ip1:port1 ip2:port2  example: cyj -tran 127.0.0.1:8888,192.168.0.100:38888")
var reflectAddr = flag.String("reflect", "", "usage: -reflect ip:port example: cyj -reflect 192.168.0.100:28888")

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	flag.Parse()
	if flag.NFlag() != 1 {
		log.Fatalln("[error]", "Use the command 'cyj -h' to get help")
		os.Exit(1)
	}
	if *listenPort != "" {
		err := CheckPort(*listenPort)
		if err != nil {
			os.Exit(1)
		}
	}
	if *monitorPort != "" {
		for _, port := range strings.Split(*monitorPort, ",") {
			err := CheckPort(port)
			if err != nil {
				os.Exit(1)
			}
		}
	}
	if *tranPort != "" {
		//fmt.Println(*tranPort)
		ipPorts := strings.Split(*tranPort, ",")
		for _, ipPort := range ipPorts {
			err := CheckIp(ipPort)
			if err != nil {
				os.Exit(1)
			}
		}
	}

	if *reflectAddr != "" {
		err := CheckIp(*reflectAddr)
		if err != nil {
			os.Exit(1)
		}
	}
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println("|               ip and port init check is ok!               |")
}

func Cyj() {
	PrintWelcome()
	switch {
	case *listenPort != "":
		err := CyjListen(*listenPort)
		if err != nil {
			os.Exit(1)
		}
	case *tranPort != "":
		addresss := strings.Split(*tranPort, ",")
		if err := CyjTran(addresss[0], addresss[1]); err != nil {
			os.Exit(1)
		}
	case *reflectAddr != "":
		err := ReflectAddress(*reflectAddr)
		if err != nil {
			os.Exit(1)
		}
		break
	case *monitorPort != "":
		ports := strings.Split(*monitorPort, ",")
		err := CyjMonitor(ports[0], ports[1])
		if err != nil {
			os.Exit(1)
		}
	default:
		break
	}
	//time.Sleep(time.Duration(3)*time.Second)
}

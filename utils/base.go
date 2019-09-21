package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 输出err
//func CheckErr(err error)  {
//	if err != nil{
//		log.Fatalln(err)
//	}
//}

// 欢迎
func PrintWelcome() {
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println("| Welcome to use cyj v0.0.1.                                |")
	fmt.Println("| Code by caihy05 at 2019-9-21 11:11:11.                    |")
	fmt.Println("| If you have some problem when you use the tool,           |")
	fmt.Println("| please send email to me 228417442@qq.com.                 |")
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println()
	time.Sleep(time.Second)
}

// 检查端口
func CheckPort(port string) error {
	portNum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln("[x]", "port should be a number")
		return err
	}
	if portNum < 20000 || portNum > 60000 {
		log.Fatalln("[x]", "port should be a number and the range is [20000,60000]")
		errMsg := errors.New("port should be a number and the range is [20000,60000]")
		return errMsg
	}
	return nil
}

func CheckIp(address string) error {
	ipAndPort := strings.Split(address, ":")
	if len(ipAndPort) != 2 {
		log.Fatalln("[x]", "address error. should be a string like [ip:port]. ")
	}
	ip := ipAndPort[0]
	port := ipAndPort[1]
	CheckPort(port)
	pattern := `^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$`
	ok, err := regexp.MatchString(pattern, ip)
	if err != nil || !ok {
		log.Fatalln("[x]", "ip error. ")
	}
	return nil
}

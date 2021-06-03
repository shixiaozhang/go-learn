package main

import (
	"flag"
	"fmt"
	"go/types"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error", err)
	}
	client.conn = conn
	// 返回对象
	return client
}

var serverIp string
var serverPort int

// ?解析命令行的格式
// ./client -ip 127.0.0.1 -port 8000
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器的ip地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8000, "设置服务器的port端口（默认是8000）")
}
func main() {
	// 命令行解析
	flag.Parse()
	// client := NewClient("127.0.0.1", 8000)
	
	client := NewClient(serverIp,serverPort)
	fmt.Println(serverIp,serverPort, client)
	if client == nil {
		fmt.Println(">>>>> 链接失败。。。")
		return 
	}
	fmt.Println(">>>>> 链接服务器成功 。。。")
	// 启动客户端业务
	select {}
}

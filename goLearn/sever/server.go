package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	// 消息广播的channel
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message 广播消息 channel 的goRoutine ，一旦有消息就发给全部在线的User
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}
func (this *Server) Handler(conn net.Conn) {
	// ... 当前链接业务
	fmt.Println("链接成功！")

	user := NewUser(conn, this)
	// // 把用户加入OnlineMap中
	// this.mapLock.Lock()
	// this.OnlineMap[user.Name] = user
	// this.mapLock.Unlock()
	// // 广播当前用户上线消息
	// this.BroadCast(user, "一上线")
	user.OnLine()
	// 监听当前用户的状态
	isLive := make(chan bool)
	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// this.BroadCast(user,"下线")
				user.Unline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			// 提取用户消息 去除‘\n’
			msg := string(buf[:n-1])
			// 将得到的消息广播
			// this.BroadCast(user,msg)
			user.DoMessage(msg)
			isLive <- true
		}
	}()
	for {
		select {
		case <-isLive:
			// 当前用户是活跃的，应该重制定时器
			// 不做任何操作，为了激活select，更新下面的定时器

		case <-time.After(time.Second * 100):
			// 已经超时了
			// 将当前的user强制关闭
			user.SendMsg("你被提了")
			// 销毁用的资源
			close(user.C)
			// 关闭链接
			conn.Close()
			// 结束go Handler

			return
		}
	}

}

// 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	// close listen socket
	defer listener.Close()
	// 启动监听Message的goroutine
	go this.ListenMessage()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
			return

			continue
		}
		//do handler
		go this.Handler(conn)
	}

}

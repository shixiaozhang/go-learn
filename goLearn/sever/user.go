package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

//创建一个用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动监听当前 user channel 消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户上线
func (this *User) OnLine() {
	// 把用户加入OnlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播当前用户上线消息
	this.server.BroadCast(this, "一上线")
}

// 用户下线,在onlinemap中删除
func (this *User) Unline() {
	// 把用户加入OnlineMap中
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	// 广播当前用户上线消息
	this.server.BroadCast(this, "下线")
}

// 返回用户消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg + "\n"))
}

//  用户发送消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询那些用户在线 消息格式 ： who
		this.server.mapLock.Lock()

		for _, user := range this.server.OnlineMap {
			onLineMap := "[" + user.Addr + "]" + user.Name + ":" + "在线呢 。。。 \n"
			this.SendMsg(onLineMap)
		}
		this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 修改用户名 ； 消息格式 ： rename｜张三

		newName := strings.Split(msg, "|")[1]
		// 判断name是否已存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名已被使用 \n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this

			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg("修改成功 \n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式： to｜占三｜消息内容
		// 1获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确,请按照\"to|张三|消息\"格式。\n")
			return

		}
		//2 根据用户名获取对方的User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户不存在 \n")
			return
		}
		// 3 获取消息内容，通过对方的User对象将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重发\n")
			return
		}
		remoteUser.SendMsg(this.Name+"对你说"+content)
	} else {
		this.server.BroadCast(this, msg)
	}
}

//创建当前 User channel的方法，一旦有消息，就直接发送给对端客户端

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

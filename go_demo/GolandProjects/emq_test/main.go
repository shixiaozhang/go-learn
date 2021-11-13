package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/eclipse/paho.mqtt.golang"
)

//func init() {
//	ff, err := os.OpenFile("./js.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	log.SetOutput(ff)
//}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {

	// 创建日志记录，flag参数定义日志属性
	mqtt.DEBUG = log.New(os.Stdout, "debug ", 0)
	mqtt.ERROR = log.New(os.Stdout, "error ", 0)

	// 建立与mqtt的连接
	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.11.241:1883").SetClientID("emqx_test_client").SetUsername("paho")
	// 设置客户端发送ping的间隔，保证客户端知道连接没有丢失
	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)

	opts.SetPingTimeout(1 * time.Second)
	// 创建客户端
	c := mqtt.NewClient(opts)
	// 创建连接
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("/test", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 发布消息
	token := c.Publish("/test", 0, false, "Hello World")
	token.Wait()	//等待服务端的确认接收

	time.Sleep(6 * time.Second)

	//test:=make(chan int)
	//<-test
	// 取消订阅
	if token := c.Unsubscribe("/test"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 断开连接
	c.Disconnect(250)
	time.Sleep(1 * time.Second)
}

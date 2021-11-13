package main

import (
	"fmt"
	"net/rpc"
	"rpc_protobuf/message"
	"time"
)

func main() {
	client,err:=rpc.DialHTTP("tcp",":8081")
	if err!=nil{
		panic(err.Error())
	}
	timeStamp:=time.Now().Unix()
	request:=message.OrderRequest{OrderId: "201907310001",TimeStamp: timeStamp}

	var response *message.OrderInfo
	err=client.Call("OrderService.GetOrderInfo",request,&response)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(*response)
}

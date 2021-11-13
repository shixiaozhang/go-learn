package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/message"
	"time"
)

func main() {
	// Dail连接
	conn,err:=grpc.Dial("localhost:8090",grpc.WithInsecure())
	if err!=nil{
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient:=message.NewOrderServiceClient(conn)

	orderRequest:=&message.OrderRequest{OrderId: "201907310002",TimeStamp: time.Now().Unix()}
	orderInfo,err:=orderServiceClient.GetOrderInfo(context.Background(),orderRequest)

	//fmt.Println(orderInfo,err)
	if orderInfo!=nil{
		fmt.Println(orderInfo.OrderId)
		fmt.Println(orderInfo.OrderName)
		fmt.Println(orderInfo.OrderStatus)
	}
}
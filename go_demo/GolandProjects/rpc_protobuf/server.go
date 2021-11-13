package main

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"rpc_protobuf/message"
	"time"
)

// 订单服务
type OrderService struct {

}

func (os *OrderService) GetOrderInfo(request message.OrderRequest,resopnse *message.OrderInfo) error {
	orderMap:=map[string]message.OrderInfo{
		"201907300001":{OrderId: "201907300001",OrderName: "衣服",OrderStatus: "已付款"},
		"201907310001":{OrderId: "201907310001",OrderName: "零食",OrderStatus: "已付款"},
		"201907310002":{OrderId: "201907310002",OrderName: "食品",OrderStatus: "未付款"},
	}
	current:=time.Now().Unix()
	if request.TimeStamp>current{
		*resopnse=message.OrderInfo{OrderId: "0",OrderName: "",OrderStatus: "订单信息异常"}
	}else {
		result:=orderMap[request.OrderId]
		if result.OrderId!=""{
			*resopnse=orderMap[request.OrderId]
		}else {
			return errors.New("server error")
		}
	}
	return nil
}

func main() {
	orderService:=new(OrderService)
	rpc.Register(orderService)
	rpc.HandleHTTP()
	listen,err:=net.Listen("tcp",":8081")
	if err!=nil{
		panic(err.Error())
	}
	http.Serve(listen,nil)
}



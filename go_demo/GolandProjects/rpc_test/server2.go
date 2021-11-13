package main

import (
	"net"
	"net/http"
	"net/rpc"
	"rpc_test/param"
)

type MathUtil2 struct {

}

func (mu MathUtil2) Add(param param.AddParam,resp *float32)error  {
	*resp=param.Args1+param.Args2
	return nil
}

func main() {
	mathUtil:=new(MathUtil2)
	// 自定义服务名称
	err:=rpc.RegisterName("MathUtil2",mathUtil)
	if err!=nil{
		panic(err.Error())
	}
	rpc.HandleHTTP()
	listen,err:=net.Listen("tcp",":8082")
	http.Serve(listen,nil)
}


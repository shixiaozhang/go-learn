package main

import (
	"fmt"
	"net/rpc"
)

//客户端逻辑实现
func main() {
	client,err:=rpc.DialHTTP("tcp","localhost:8081")
	if err!=nil{
		panic(err.Error())
	}
	var req float32		// 请求值
	req=4

	//同步的调用方式
	//var resp*float32	// 返回值
	//err = client.Call("MathUtil.CalculateCircleArea",req,&resp)
	//if err!=nil{
	//	panic(err.Error())
	//}
	//fmt.Println(*resp)

	// 异步的调用方式
	var respSync *float32
	syncCall:=client.Go("MathUtil.CalculateCircleArea",req,&respSync,nil)
	fmt.Println(respSync)
	replayDone:=<-syncCall.Done
	fmt.Println(replayDone)
	fmt.Println(*respSync)

}
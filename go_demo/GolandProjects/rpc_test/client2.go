package main

import (
	"fmt"
	"net/rpc"
	"rpc_test/param"
)

func main() {
	client,err:=rpc.DialHTTP("tcp",":8082")
	if err!=nil{
		panic(err.Error())
	}
	var result *float32
	addParam:=&param.AddParam{Args1: 1.2,Args2: 2.3}
	err=client.Call("MathUtil2.Add",addParam,&result)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println("计算结果：",*result)

}

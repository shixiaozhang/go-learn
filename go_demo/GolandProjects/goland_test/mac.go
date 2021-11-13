package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	Mac2()
	fmt.Println("NODE_ID:", os.Getenv("NODE_ID"))
}

func Mac2(){
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	//for _, inter := range interfaces {
	//fmt.Println(inter.Name)
	inter := interfaces[0]
	mac := inter.HardwareAddr.String() //获取本机MAC地址
	fmt.Println("MAC = ", mac)
	err = os.Setenv("NODE_ID", mac)
	if err != nil {
		fmt.Println("ERROR:NODE_ID SET----", err.Error())
	}

	//}
}
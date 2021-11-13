package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// 单线程TCP端口扫描器
//func main() {
//
//	for i := 15660; i <15680 ; i++ {
//		address:=fmt.Sprintf("127.0.0.1:%d",i)
//		conn,err:=net.Dial("tcp",address)
//		if err!=nil{
//			fmt.Printf("%s关闭了\n",address)
//			continue
//		}
//		conn.Close()
//		fmt.Printf("%s打开了\n",address)
//
//	}
//}

// 多线程TCP端口扫描器
func main() {
	// 获取当前时间
	start:=time.Now()
	// 计数等待
	var wg sync.WaitGroup
	for i := 15660; i <15680 ; i++ {
		wg.Add(1)
		// 定义匿名函数并直接传参运行
		go func(j int) {
			defer wg.Done()
			address:=fmt.Sprintf("127.0.0.1:%d",j)
			conn,err:=net.Dial("tcp",address)
			if err!=nil{
				fmt.Printf("%s关闭了\n",address)
				return
			}
			conn.Close()
			fmt.Printf("%s打开了\n",address)
		}(i)
	}
	wg.Wait()
	// 计算相隔时间	纳秒转秒
	elapsed:=time.Since(start)/1e9
	fmt.Printf("\n\n%d seconds",elapsed)

}

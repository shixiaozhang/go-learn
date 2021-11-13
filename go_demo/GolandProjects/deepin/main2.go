package main

import (
	"fmt"
	"net"
	"sync"
)

// work 工作池
func worker(ports chan int,wg *sync.WaitGroup,results chan int)  {
	for p:=range ports{
		//fmt.Println(p)
		address:=fmt.Sprintf("127.0.0.1:%d",p)
		conn,err:=net.Dial("tcp",address)
		if err!=nil{
			//fmt.Printf("%s关闭了\n",p)
			results<-0
			wg.Done()
			continue
		}
		conn.Close()
		//fmt.Printf("%s打开了\n",p)
		results<-p
		wg.Done()
	}

}

func main() {
	ports:=make(chan int,100)
	results:=make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports,&wg,results)
	}
	for i := 1; i <1024 ; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)

}

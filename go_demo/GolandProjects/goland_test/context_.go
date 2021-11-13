package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethoing(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("finish doing something")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func main() {
	// 创建空context的两种方法
	ctx := context.Background() // 返回一个空的context，不能被cancel，kv为空
	//todoCtx:=context.TODO()		// 和Background类似，当你不确定的时候使用

	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	//ctx, cancel := context.WithDeadline(ctx,time.Now())
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	doSomethoing(ctx)

}

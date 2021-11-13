package main

import "fmt"



func main() {
	// 递归、阶乘
	var i = 10
	fmt.Printf("%d的阶乘是%d\n",i,Factorial(uint64(i)))

	// 接口数据类型
	var phone Phone
	// 创建指针变量
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()


}


func Factorial(n uint64) (result uint64) {
	if n>0{
		result=n*Factorial(n-1)
		return result
	}
	return 1
}

// 接口类型 多态
type Phone interface {
	call()
}

type NokiaPhone struct {
	high int
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {

}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}





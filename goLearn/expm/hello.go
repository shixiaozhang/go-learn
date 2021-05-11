package hello

import (
	"fmt"
	"time"
)

var gA int = 100
var gB = 200

func Hello() {
	fmt.Println("hello go!")
	time.Sleep(1 * time.Second)
	fmt.Println("hello go ...end!")
	// ? 直接声明一个变量，默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of b = %T \n", a)
	// ? 声明变量的时候直接初始化一个值；
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b =%T\n", b)

	var bb string = "bd"
	fmt.Printf("bb = %s, type of bb =%T\n", bb, bb)
	// ? 初始化省去类型声明；自动推导类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c =%T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	// ? 省去var 关键字 直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type if e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf(" type of f =%T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g =%T\n", g)

	fmt.Println("gA = ", gA, "gB = ", gB)

	//? 一次声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)
	var kk, ll = 100, "abcd"
	fmt.Println("kk = ", kk, "ll = ", ll)

	// ?多行声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}

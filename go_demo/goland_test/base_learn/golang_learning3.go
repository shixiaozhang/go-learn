package main

import "fmt"

func main() {
	// 切片初始化（动态数组）
	s:= [] int {1,2,3}
	// 初始化数组
	balance:=[5]float32{100.0,2.0,3.6,8.5,6.7}
	// 切片截取
	n:= balance[1:4]
	// make 初始化切片
	var m =make([]int,3,5)
	// cap 返回切片可达的最长长度
	fmt.Printf("len=%d cap=%d slice=%v\n",len(m),cap(m),m)
	fmt.Println(s,n)
	var numbers [] int
	if numbers==nil{
		fmt.Println("切片是空的")
	}
	// append
	s=append(s,3,5)
	fmt.Println(s,cap(s))
	// 创建两倍容量的切片
	s2:=make([]int,len(s),(cap(s))*2)
	copy(s2,s)
	fmt.Println("1",s2,cap(s2))

	// range 返回key-value对 数组、切片、通道、集合
	for k,v :=range s{
		fmt.Println(k,v)
		break
	}

	// range 枚举Unicode字符串
	for i,c :=range "go"{
		fmt.Println(i,c)
	}

	// 集合操作（python中的字典）
	//country:=make(map[string]string,10)
	country:=map[string]string{
		"one":"nanjing",
		"two":"shanghai",
	}
	country["Japan"]="东京"
	// 删除
	delete(country,"Japan")
	// 判断元素是否在集合中
	capital,ok:=country["American"]
	if ok{
		fmt.Println("American的首都是",capital)
	}else {
		fmt.Println("American首都不存在")
	}



}




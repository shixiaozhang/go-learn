package main

import "fmt"

func main()  {
	fmt.Println("1111")
	var stock=123
	enddate:="2021-11-13"

	fmt.Println(stock,enddate)
	url :="Code=%v ednddate=%s"
	fmt.Sprintf(url,stock,enddate)
	fmt.Println(stock)
	var b int
	fmt.Printf("type of b = %T\n",b)
	const length = 10
	const width = 5
	area :=length* width
	fmt.Printf("面积为： %v\n",area)
	var list [10]int
	for  i:=0;i<len(list) ; i++{
		fmt.Println(list[i],i)
	}
	for i, i2 := range list {
		fmt.Println(i,i2)
		list[i]=i
	}
	fmt.Println(list)
	var list2 =make([]int,10,20)
	list2=append(list2,11)
	fmt.Printf("%v \n",list2)
	fmt.Println(len(list2),cap(list2))
	fmt.Println(len(list),cap(list))
	 listArr :=[2]int{11,21}
	fmt.Printf("%v \n %T",listArr,listArr)
	fmt.Println(listArr)

	//1.不同类型的切片无法复制
	//2.如果s1的长度大于s2的长度，将s2中对应位置上的值替换s1中对应位置的值
	//3.如果s1的长度小于s2的长度，多余的将不做替换
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s3 := []int{6, 7, 8, 9}
	copy(s1, s2)
	fmt.Println(s1) //[4 5 3]
	copy(s2, s3)
	fmt.Println(s2) //[6 7]
	// 函数 copy 在两个 slice 间复制数据，复制⻓度以 len 小的为准，两个 slice 指向同⼀底层数组。直接对应位置覆盖。
}
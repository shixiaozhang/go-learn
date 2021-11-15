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
	mapObj:=map[string]string{}
	var mapObj2 map[string]string
	if mapObj == nil {
		fmt.Println("空map")
	}else {
		fmt.Println(mapObj,len(mapObj))
	}
	if mapObj2 == nil {
		fmt.Println("kong map")
	}
	mapObj3 :=make(map[string]string,2)
	if mapObj3 ==nil {
		fmt.Println("kong map")
	}else {
		fmt.Println(mapObj3,len(mapObj3))
	}
	mapObj["小张"]="张"
	mapObj["小吴"]="吴"
	fmt.Println(mapObj)
	mapObj3["张"]="小张"
	mapObj3["吴"]="小吴"
	mapObj3["xx"]="小xx"
	fmt.Println(mapObj3,len(mapObj3))
	mapObj4:=make(map[int]string)
	mapObj4[1]="11"
	fmt.Println(mapObj4)
	mapObj5:=map[string]string{
		"xiaozhang":"zhang",
		"xiaowu":"wu",
	}
	fmt.Println(mapObj5)
	// 删除
	delete(mapObj3,"xx")

	// 判断元素是否在集合中
isHaveZ,ok:=mapObj["小张"]
	if ok {
		fmt.Println("have a zhang,is ",isHaveZ)
	}else {
		fmt.Println("not have zhang")
	}
	fmt.Println(mapObj3)

	for key, value := range mapObj {
		fmt.Println(key,value)
	}

	for k, v := range "collection" {
		fmt.Println(k,v)
	}

	var book Book
	book.id=1
	book.context="687648618764871"

	context:=changeBook(book)
	fmt.Println(context)

	fmt.Printf("%v\n", book) //{1  687648618764871}
	var book1 Book
	book1.id=2
	book1.context="687648618764871"
	changeBookPointer(&book1)
	fmt.Printf("%v\n", book1) //{2 wu 687648618764871}

	//	类
	hero :=Hero{Name: "zhang",age: 18,Sex: "nan"}
	name:=hero.getName()
		fmt.Println(name)
	hero.showInfo()
	hero.setName("xiaozhang")
	fmt.Println(hero)
}
//自定义数据类型
type myint int
//结构体
type  Book struct {
	id myint
	name string
	context string
}

func changeBook(book Book) string {

	book.name="zhang"
	return book.context
}
func changeBookPointer(book *Book) string  {
	book.name="wu"
	return book.context
}
//类
type Hero struct {
	Name string
	Sex string
	age int
}

func (this Hero) showInfo() {
	fmt.Println(this.Name)
	fmt.Println(this.Sex)
	fmt.Println(this.age)
}

func (this Hero) getName() string {
	return this.Name
}

func (this Hero) setName(name string) {
	this.Name=name
}
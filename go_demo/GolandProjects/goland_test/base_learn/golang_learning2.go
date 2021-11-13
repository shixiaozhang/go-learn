package main

import "fmt"

func main(){
	// 声明一个数组
	var balance [10] int
	// 初始化一个数组
	//balance2:=[5]float32{100.0,2.0,3.6,8.5,6.7}
	// 长度不确定时用[...]
	var balance3=[...]float32{100.0,2.0,3.6,8.5,6.7}
	// 指定索引初始化
	balance4:=[5]float32{1:2.0,3:7.0}
	balance4[4]=6.9
	// 访问数组
	var nn float32=balance3[2]
	fmt.Println(nn)

	var i,j int
	for i=0;i<10;i++{
		balance[i]= i + 100
	}
	for j=0;j<10;j++{
		fmt.Printf("Element[%d] = %d\n", j, balance[j] )
	}

	// 结构体
	fmt.Println(Books{"隐私计算","黑云",4396})
	fmt.Println(Books{title: "区块链"})

	// 声明结构体变量
	var Book1 Books
	Book1.title="NLP"
	Book1.author="黑云"
	Book1.book_id=6789
	fmt.Printf("Book1 title : %s\n",Book1.title)

	// 结构体作为函数参数
	printBook(Book1)
	// 调用结构体的方法
	Book1.Show()
	Book1.SetName("自然语言处理")
	Book1.Show()

	// 结构体指针
	var struct_pointer *Books
	struct_pointer=&Book1
	fmt.Println(struct_pointer.title)

	// 定义子类对象
	SuperBook1:=SuperBooks{Books{"11","22",33},44}
	fmt.Println(SuperBook1)
}

type Books struct {
	title string
	author string
	book_id int
}

type SuperBooks struct {
	Books	// 继承Books类的方法
	level int
}


// 普通函数接受结构体参数
func printBook( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
}

// 结构体类的方法，表示对原结构体的copy
func (this Books) Show()  {
	fmt.Println("书的标题：",this.title)
}

// 指向原结构体地址，表示同一个结构体，可修改
func (this *Books) SetName(newName string)  {
	this.title=newName

}

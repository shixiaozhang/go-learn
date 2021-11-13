// package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
package main

// 这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数
import "fmt"

// main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
func main() {
	fmt.Println("hello world")

	// 定义变量 字符串拼接(格式化字符串)
	var stockcode = 123
	// :=用于声明变量
	enddate := "2020-12-31"
	var url = "Code=%dendDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	// 打印数据的类型
	var b int
	fmt.Printf("type of b = %T\n",b)

	// 定义常量
	const length = 10
	const width = 5
	area := length * width
	fmt.Printf("面积为：%d\n", area)

	// 添加一个关键字iota，初始为0，每行都会增加一个
	const (
		beijing=iota
		shanghai
		nanjing
	)
	fmt.Println(beijing,shanghai,nanjing)

	// 逻辑运算
	if stockcode > width {
		fmt.Println(length, width)
	} else {
		fmt.Println(length, width)
	}

	// & 返回变量存储地址 *指针变量
	var a = 4
	var ptr *int

	ptr = &a
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// defer关键字，函数退出前执行，多个defer符合栈的先进后出
	// defer关键字比return还晚执行
	defer fmt.Println("defer测试")


	// 循环语句
	for true {
		fmt.Println("无限循环")
		break
	}

	// 定义mxa函数
	var m = 100
	var n = 200
	var ret int
	ret = max(m, n)
	fmt.Printf( "最大值是 : %d\n", ret )

}


/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 导包要导绝对路径，包内包含init函数，其他函数（类）首字母大写为公有方法、首字母小写为私有方法
// 导包细节详见routers

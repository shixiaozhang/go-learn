package main 
//?声明 这个文件也就是这个包的类型
//?导入语句，必须用双引号
import (
	"fmt"
	"math"
	"strings"
)

//? import fm "fmt" //? alias3//? 把fmt 重命名为 fm
//?在函数外边只能放 标识符 （变量、常量、函数 、类型 ）变量的；不能放运行的语句

var name string

//?Iz intleixing
type Iz int

var a Iz = 5
// var all string

const chang = "1"
const ( //? 这种情况三个fo 都是100
	fo = 100
	fo1
	fo2
)
const ( //? 这三值 是 0 1 2
	f1 = iota
	f2
	f3
)
const (
	d1 = iota //? 0
	_         //?跳过
	d3        //?2
)

const (
	e1 = iota //?0
	e2 = 155  //?155
	e3        //? 155
	e4 = iota //?3
	e5        //?4
	e6        //?5
)

//? 函数外的变量可以申明，但是不赋值，也可以不使用
//? 但是函数内的变量必须使用，和赋值
//? ?以大写字母开头的变量，相当于暴露，可被外部的报使用
//? ?先引入这个包，再用这个包 。 变量名使用
func foo() (string, string) {
	return "111", "wwww"
}
func main() {
	//? 这个包的入口函数，必须有;main函数没有参数和返回值
	//? 变量必须先声明在使用,声明的变量必须使用,字符串要用双引号；
	//? ?变量声明后没赋值，不能再函数外赋值；必须在函数内
	name = "xiaozhang"
	hh := 'j'
	hh2 := "jhkashdkjh"
	//? 变量声明尽量用双引号，单引号只用，定义一个字节的字符串

	x, _ := foo() //? 匿名变量 _

	fmt.Println("hello")
	fmt.Printf("%s \n", "aaa")
	fmt.Printf("%s rwer  \n", "aaa")
	fmt.Printf("%d \n", a)
	fmt.Println(name)
	fmt.Println(hh, hh2, x, chang)
	fmt.Println(fo, fo1, fo2)
	fmt.Println(f1, f2, f3)
	fmt.Println(d1, d3)
	fmt.Println(e1, e2, e3, e4, e5, e6)

	var b int = 077
	fmt.Printf("%o \n", b)

	ss := `第一行
			第二行
			第三行`

	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	s1 := "字符串"
	fmt.Println(ss)

	length := len(s1)

	str := "我是一个字符串！"
	fmt.Printf("%v \n", length+b)

	str2 := fmt.Sprintf("%s+%s \n", s1, str) //?Sprintf是拼接后返回拼接的值
	fmt.Printf("str2: %s \n", str2)

	fmt.Printf("%s%s \n", s1, str)          //? Printf 是打印
	fmt.Println(strings.Split(str, ""))     //?字符串分割
	fmt.Println(strings.Contains(str, "一")) //? 字符串是否包含，返回Boolean
	str4 := "字符.txt"
	fmt.Println(strings.HasPrefix(str4, "字符"))
	fmt.Println(strings.HasSuffix(str4, "字符"))
	str5 := "abcdfc"
	fmt.Println(strings.Index(str5, "c"))     //? ?第一次出现的位置
	fmt.Println(strings.LastIndex(str5, "c")) //? 最后一次出现的位置

	str6 := strings.Split(str, "")
	str7 := strings.Join(str6, "+")
	fmt.Println(str7)
	for i := 0; i < len(str); i++ { //?byte
		fmt.Printf("%v (%c)", str[i], str[i])
	}
	// fmt.Println("\n 遍历数组 \n")
	for _, v := range str {
		fmt.Printf("%v (%c)", v, v)
	}
	// fmt.Println("\n 遍历数组 \n")

	s12 := "big"
	//? 强制类型转换
	byteS1 := []byte(s12)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1)) //? pig

	s22 := "白萝卜"
	runeS2 := []rune(s22)
	runeS2[0] = '红'
	fmt.Println(string(runeS2)) //?	红萝卜
	var a1, b1 = 3, 4
	var c1 int
	//? math.Sqrt()接收的参数是float64类型，需要强制转换
	c1 = int(math.Sqrt(float64(a1*a1 + b1*b1)))
	fmt.Println(c1) //? 5
	//? if else
	ifAge := 19
	//? 基础
	if ifAge > 18 {
		fmt.Println("dada")
	} else if ifAge > 10 {
		fmt.Println("aaaa")
	} else {
		fmt.Println("ccc")
	}
	//?带作用域的if
	if ifAge := 19; ifAge > 18 {
		fmt.Println("带作用域的if")
	} else {
		fmt.Println("带作用域的if2222")
	}

	//? for循环

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//? for循环变种1
	var forI int = 5
	fmt.Println(forI)

	for ; forI < 10; forI++ {
		fmt.Println(forI)
	}
	//? for循环变种2

	var forI2 int = 5
	fmt.Println(forI2)

	for forI2 < 10 {
		fmt.Println(forI2)
		forI2++
	}
	//? for range(键值循环) 返回索引和值。
	rangeVal := "nihao战三"

	for i, v := range rangeVal {
		//? i 是索引 v是值
		//? 英文字母站一个索引，也就是一个字节；中文字符站三个索引，也就是三个字符
		fmt.Printf("%d %c \n", i, v)
		//? 0 n
		//? 1 i
		//? 2 h
		//? 3 a
		//? 4 o
		//? 5 战
		//? 8 三
	}
	fmt.Println("break")
	//?  跳出
	for i := 0; i < 10; i++ {

		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("continue")
	//?  跳过
	for i := 0; i < 10; i++ {

		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	//?   333
	finger := 3
	switch finger {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}
	//?  222
	switch {
	case finger > 3:
		fmt.Println("qwe")
	default:
		fmt.Println("moren")
	}
	//?   111
	switch i := 2; i {
	case 1, 2, 3:
		fmt.Println("123")
	case 4:
		fmt.Println("4")
		fallthrough
	default:
		fmt.Println("moren")
	}
	//?   goto跳出指定breakTag

	for i := 0; i < 10; i++ {
		for k := 0; k < 10; k++ {
			if k == 2 {
				goto breakTag
			}
		}
	}
breakTag:
	fmt.Println("goto")

	//?  break跳出制定for
BBB:
	for i := 0; i < 10; i++ {
	AAA:
		for k := 0; k < 10; k++ {
			if k == 2 {
				fmt.Println("/v /v", i, k)
				break AAA
			}
		}
		if i == 3 {
			fmt.Println("/v", i)
			break BBB
		}

	}

	var intArr [3]int

	fmt.Println(intArr)    //?  [0 0 0]
	fmt.Println(intArr[2]) //?  0

	//?   初始化数组时可以使用初始化列表来设置数组元素的值。
	var testArr [3]int
	var numberArr = [3]int{1, 2}
	var cityArr = [3]string{"北京", "shanghai", "shenzhen"}
	fmt.Println(testArr)   //?  [0 0 0]
	fmt.Println(numberArr) //?  [1 2 0]
	fmt.Println(cityArr)   //?  [北京 shanghai shenzhen]

	//?   以让编译器根据初始值的个数自行推断数组的长度，
	var cccArr = [...]int{1, 2}
	fmt.Println(cccArr) //[1 2]
	//?   指定索引值的方式来初始化数组，
	aaaa := [...]int{1: 1, 3: 5}
	fmt.Println(aaaa)
	//?  遍历数组
	for i := 0; i < len(cityArr); i++ {
		fmt.Println(cityArr[i])
	}
	for index, value := range cityArr {
		fmt.Println(index, value)
	}
	//?  二维数组的定义
	ccccc := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(ccccc)
	for _, v1 := range ccccc {
		for _, v2 := range v1 {
			fmt.Printf("%s \t", v2)
		}
		fmt.Println()
	}
	//?   多维数组只有第一层可以使用...
	avcsscsc := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(avcsscsc)
	modifyArray(testArr)
	//? 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。7
	fmt.Println(testArr)
	//? 	数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
	//? [n]*T表示指针数组，*[n]T表示数组指针 。
}

//?  数组是值类型

func modifyArray(x [3]int) {
	x[0] = 100
}

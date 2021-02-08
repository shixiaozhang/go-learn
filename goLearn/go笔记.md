<!--

 * @Author: your name
 * @Date: 2021-01-20 22:53:02
 * @LastEditTime: 2021-01-25 22:33:12
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \src\github.dafu.com\goLearn\go笔记.md
-->

#  安装Go语言开发工具包

在此之前请先设置GOPROXY，打开终端执行以下命令：

    配置这些之前要先配置环境变量：
    GOPATH:项目文件目录
    path： 项目安装目录
    go安装推荐安装在c盘，不然老是出问题；


```go
go env -w GOPROXY=https://goproxy.cn,direct
```

   Ctrl+Shift+P ===>  Go:Install/Update Tools ===> 全选17项（要挂vpn）




# 配置代码片段快捷键


```go
Ctrl/Command+Shift+P,按下图输入>snippets

然后在弹出的窗口点击选择go选项;在go.json
```
   （ 无所谓）


```go
    “这里放个名字”:{
        "prefix": "这个是快捷键",
        "body": "这里是按快捷键插入的代码片段",
        "description": "这里放提示信息的描述"
    }
```



# 编译

main.go : go项目的入口文件

go build  编译
go build  /xxx/xxx/xxx 编译某个地方的go ； src目录下开始
go build -o name.exe 编译后文件的名字
# 执行 
go run main.go 执行某个文件
go install   先编译go文件 ，再把编译好的文件放入bin文件夹下


# 跨平台编译

默认我们go build的可执行文件都是当前操作系统可执行的文件，如果我想在windows下编译一个linux下可执行文件，那需要怎么做呢？

只需要指定目标操作系统的平台和处理器架构即可：

    SET CGO_ENABLED=0  // 禁用CGO
    SET GOOS=linux  // 目标平台是linux
    SET GOARCH=amd64  // 目标处理器架构是amd64

使用了cgo的代码是不支持跨平台编译的

然后再执行go build命令，得到的就是能够在Linux平台运行的可执行文件了。




Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
    
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

Linux 下编译 Mac 和 Windows 平台64位可执行程序：

    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
    
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

Windows下编译Mac平台64位可执行程序：

    SET CGO_ENABLED=0
    SET GOOS=darwin
    SET GOARCH=amd64
    go build

# Println 与 Printf 的区别，以及 Printf 的详细用法：

Println :可以打印出字符串，和变量

```
 f.Println("hello","world","hello","world")
 
 f.Println("hello","world","hello","world")
 
//输出
 hello world hello world
 
 hello world hello world
```

在同一输出函数中输出多项的时候，hello和world中是存在空格的；在不同输出函数之间会换行

Print :在同一个输出函数中处处多项的时候，hello和world中不存在空格,在不同输出函数之间，不换行



```
 f.Println("hello","world","hello","world")
 
 f.Println("hello","world","hello","world")
 
 //输出
 
 helloworldhelloworldhelloworldhelloworld
 
```

Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量，不可以输出整形变量和整形

```
   a := 10
 
   b := 20
 
   c := "hello"
 
   f.Printf("a=%d,b=%d",a,b)
 
   f.Printf("c=%s",c)

//输出

a=10,b=20c=hello

可以对参数进行格式化输出，在不同输出函数中是不换行的
```

## Printf 格式化输出

## 通用占位符：

###  **v** 值的默认格式；v 就是后面的值 ，什么值都可以

* %v      打印后面值，不管其他
* %+v    添加字段名(如结构体)
* %#v    相应值的Go语法表示 ，如果是string 会加上双引号打印 
* %T      相应值的类型的Go语法表示
* %%     字面上的百分号，并非值的占位符

## **布尔值：**

* %t  true 或 false

## 整数值：

* %b  二进制表示
* %c  相应Unicode码点所表示的字符
* %d  十进制表示
* %o  八进制表示
* %q  单引号围绕的字符字面值，由Go语法安全地转义
* %x  十六进制表示，字母形式为小写 a-f
* %X  十六进制表示，字母形式为大写 A-F
* %U  Unicode格式：U+1234，等同于 "U+%04X"

## 浮点数及复数：

* %b  无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat中的 'b' 转换格式一致。例如 -123456p-78
* %e  科学计数法，例如 -1234.456e+78
* %E  科学计数法，例如 -1234.456E+78
* %f  有小数点而无指数，例如 123.456
* %g  根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
* %G  根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出

## 字符串和bytes的slice表示：

* %s 字符串或切片的无解译字节
* %q 双引号围绕的字符串，由Go语法安全地转义
* %x 十六进制，小写字母，每字节两个字符
* %X 十六进制，大写字母，每字节两个字符



## 指针：

* %p 十六进制表示，前缀 0x
  这里没有 'u' 标记。若整数为无符号类型，他们就会被打印成无符号的。类似地，这里也不需要指定操作数的大小（int8，int64）。

## **对于％ｖ来说默认的格式是：**

* bool: %t
* int, int8 etc.: %d
* uint, uint8 etc.: %d, %x if printed with %#v
* float32, complex64, etc: %g
* string: %s
* chan: %p
* pointer: %p

# 变量和常量：

标识符：只能以字母和_ 开头

# 变量声明：变量必须先声明在使用,声明的变量必须使用,字符串要用双引号；变量声明后没赋值，不能再函数外赋值；必须在函数内


    var 变量名 类型 
    
    var name string
    var age int
    var isOk bool

 ## 批量声明：

        var (
            a string
            b int
            c bool
            d float32
        )

## 变量的初始化

var 变量名 类型 = 表达式

    var name string = "Q1mi"
    var age int = 18

或者一次初始化多个变量

   var name,age="小站",15


## 类型推导：

有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

    var name = "Q1mi"
    var age = 18

## 短变量声明：在函数内部，可以使用更简略的 := 方式声明并初始化变量


        // 全局变量m
        var m = 100
    
        func main() {
            n := 10
            m := 200 // 此处声明局部变量m
            fmt.Println(m, n)
        }

## 匿名变量

在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示，
匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明

```
    func foo() (int, string) {
        return 10, "Q1mi"
    }
    func main() {
        x, _ := foo()
        _, y := foo()
        fmt.Println("x=", x)
        fmt.Println("y=", y)
    }
```

# 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把`var`换成了`const`，**常量在定义的时候必须赋值**。

```
const pi = 3.1415
const e = 2.7182
```

多个常量也可以一起声明：

```go
const (
    pi = 3.1415
    e = 2.7182
)
```

**const同时声明多个常量时，如果省略了值则表示和上面一行的值相同**。 例如：

```go
const (
    n1 = 100
    n2
    n3
)
```

## **上面示例中，常量`n1`、`n2`、`n3`的值都是100。**但是 iota 就不一样




# iota : iota是go语言的常量计数器，只能在常量的表达式中使用



**iota 在 const 关键字出现时将被 重置 为  0 **。**const 中每新增一行常量声明将使iota计数一次**  (iota可理解为const语句块中的行索引)。 使用iota能简化定义，**在定义枚举时很有用。**

## const 中每新增一行常量声明将使iota计数一次

    const (
            n1 = iota //0
            n2        //1
            n3        //2
            n4        //3
        )

使用_跳过某些值

    const (
            n1 = iota //0
            n2        //1
            _
            n4        //3
        )

iota声明中间插队

    const (
            n1 = iota //0
            n2 = 100  //100
            n3 = iota //2
            n4        //3
          )
    const n5 = iota //0

多个iota定义在一行

        const (
                a, b = iota + 1, iota + 2 //1,2
                c, d                      //2,3
                e, f                      //3,4
            )

定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）

        const (
                _  = iota
                KB = 1 << (10 * iota)
                MB = 1 << (10 * iota)
                GB = 1 << (10 * iota)
                TB = 1 << (10 * iota)
                PB = 1 << (10 * iota)
            )


# 注意：

## 1、函数外的每个语句都必须以关键字开始（var、const、func等）

## 2、:= 不能使用在函数外。

## 3、_ 多用于占位，表示忽略值。





# Go 程序的一般结构

* 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
* 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。
* 如果当前包是 main 包，则定义 main 函数。
* 然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。

```
package main

import (
   "fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

func main() {
   var a int
   Func1()
   // ...
   fmt.Println(a)
}

func (t T) Method1() {
   //...
}

func Func1() { // exported function Func1
   //...
}

```

# Go 程序的执行（程序启动）顺序如下：

1、按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
2、如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
3、然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
4、在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。

##  类型转换

**在必要以及可行的情况下，一个类型的值可以被转换成另一种类型的值。由于 Go 语言不存在隐式类型转换，因此所有的转换都必须显式说明，就像调用一个函数一样（类型在这里的作用可以看作是一种函数）：**

```
valueOfTypeB = typeB(valueOfTypeA)

类型 B 的值 = 类型 B (类型 A 的值)

a := 5.0
b := int(a)
```

但这只能在定义正确的情况下转换成功，例如**从一个取值范围较小的类型转换到一个取值范围较大的类型**（例如将 int16 转换为 int32）。**当从一个取值范围较大的转换到取值范围较小的类型时（例如将 int32 转换为 int16 或将 float32 转换为 int），会发生精度丢失（截断）的情况**。当编译器捕捉到非法的类型转换时会引发编译时错误，否则将引发运行时错误。

具有相同底层类型的变量之间可以相互转换

```
var a IZ = 5
c := int(a)
d := IZ(c)
```

## Go 命名规范

**名称不需要指出自己所属的包**，因为在调用的时候会使用包名作为限定符。**返回某个对象的函数或方法的名称一般都是使用名词**，没有 Get... 之类的字符，如果是用于修改某个对象，则使用 SetName。有必须要的话可以使用大小写混合的方式，如 MixedCaps 或 mixedCaps，而不是使用下划线来分割多个名称。

# 可见性规则

当标识符（包括常量、变量、类型、函数名、结构字段等等）以**一个大写字母开头**，如：Group1，那么使用这种形式的标识符的**对象就可以被外部包的代码所使用**（客户端程序需要先导入这个包），这被称为**导出**（像 **面向对象语言中的 public**）；标识符如果**以小写字母开头，则对包外是不可见的**，但是他们**在整个包的内部是可见并且可用的**（像 **面向对象语言中的 private** ）。

（大写字母可以使用任何 Unicode 编码的字符，比如希腊文，不仅仅是 ASCII 码中的大写字母）。

因此，在**导入一个外部包后，能够且 *只能*  够访问该包中导出的对象**。

假设在包 pack1 中我们有一个变量或函数叫做 Thing（以 T 开头，所以它能够被导出），那么在当前包中导入 pack1 包，Thing 就可以像面向对象语言那样使用点标记来调用：**pack1.Thing**（pack1 在这里是不可以省略的）。

因此包也可以作为命名空间使用，帮助避免命名冲突（名称冲突）：两个包中的**同名变量的区别在于他们的包名**，例如 pack1.Thing 和 pack2.Thing。

你可以**通过使用包的别名来解决包名之间的名称冲突**，或者说根据你的个人喜好对**包名进行重新设置**，如：**import fm "fmt"**。下面的代码展示了如何使用包的别名：



```
// alias.go
package main

import fm "fmt" // alias3// 把fmt 重命名为 fm

func main() {
   fm.Println("hello, world")


```



# Go语言基础之基本数据类型

可以包含数据的变量（或常量），可以使用不同的数据类型或类型来保存数据。使用 **var 声明的变量的值会自动初始化为该类型的零值**。类型定义了某个变量的值的集合与可对其进行操作的集合。

Go语言中有丰富的数据类型，除了基本的**整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）**等。Go 语言的基本类型和其他语言大同小异。

基本类型，如：**int、float、bool、string**；结构化的（复合的），如：**array、slice、struct、func、map、channel**；只描述类型的行为的，如：**interface** ;值得注意的是，Go 语言中**不存在类型继承**。

结构化的类型没有真正的值，它使用 **nil** 作为默认值（**在 Objective-C 中是 nil，在 Java 中是 null，在 C 和 C++ 中是 NULL 或 0**）

函数也可以是一个确定的类型，就是以函数作为返回类型。这种类型的声明要写在函数名和可选的参数列表之后

```

func FunctionName (a typea, b typeb) typeFunc

可以在函数体中的某处返回使用类型为 typeFunc 的变量 var：

return var

//--------

一个函数可以拥有多返回值，返回类型之间需要使用逗号分割，并使用小括号 () 将它们括起来

func FunctionName (a typea, b typeb) (t1 type1, t2 type2)
返回的形式：
return var1, var2
```

使用 type 关键字可以定义你自己的类型，你可能想要定义一个结构体 (第 10 章)，但是也可以定义一个已经存在的类型的别名

```

type IZ int

使用下面的方式声明变量

var a IZ = 5

有多个类型需要定义，可以使用因式分解关键字的方式，例如：

type (
   IZ int
   FZ float64
   STR string
)
```

**这里并不是真正意义上的别名，因为使用这种方法定义之后的类型可以拥有更多的特性，且在类型转换时必须显式转换。**

**每个值都必须在经过编译后属于某个类型（编译器必须能够推断出所有值的类型），因为 Go 语言是一种静态类型语言。**

# 整型
整型分为以下两个大类： 
* 按长度分为：int8、int16、int32、int64 

* 对应的无符号整型：uint8、uint16、uint32、uint64

  

  ### 特殊整型

| uint    | 32位操作系统上就是`uint32`，64位操作系统上就是`uint64` |
| ------- | ------------------------------------------------------ |
| int     | 32位操作系统上就是`int32`，64位操作系统上就是`int64`   |
| uintptr | 无符号整型，用于存放一个指针                           |



**注意事项** 获取对象的长度的内建`len()`函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用`int`来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用`int`和 `uint`。

# 数字字面量语法

借助fmt函数来将一个整数以不同进制形式展示。

```go
package main
 
import "fmt"
 
func main(){
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10   %d 表示十进制
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
 
	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77 占位符 %o 表示八进制 是o 不是 零0
 
	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff 占位符%x表示十六进制
	fmt.Printf("%X \n", c)  // FF
```

# 浮点型

## **Go语言支持两种浮点型数：`float32`和`float64`**；默认的小数 是float64类型

**`float32` 的浮点数的最大范围约为 `3.4e38`，可以使用常量定义：`math.MaxFloat32`。 `float64` 的浮点数的最大范围约为 `1.8e308`，可以使用一个常量定义：`math.MaxFloat64`。**

```
package main
import (
        "fmt"
        "math"
)
func main() {
        var fl6 ：=1.2222
        fl32 :=float32(1.333)
        
        fmt.Printf("%T\n", fl6 )//float64 %T 打印类型
        fmt.Printf("%T\n", fl32 )//float32 %T 打印类型
        fmt.Printf("%f\n", math.Pi)
        fmt.Printf("%.2f\n", math.Pi)
}
```

# 布尔值

## 默认是 false 类型

**Go语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true（真）`和`false（假）`两个值。**

# 字符串

**Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用`UTF-8`编码。 字符串的值为`双引号(")`中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：**

```go
s1 := "hello"
s2 := "你好"
```

### 字符串转义符

Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义符 |                含义                |
| :----: | :--------------------------------: |
|  `\r`  |         回车符（返回行首）         |
|  `\n`  | 换行符（直接跳到下一行的同列位置） |
|  `\t`  |               制表符               |
|  `\'`  |               单引号               |
|  `\"`  |               双引号               |
|  `\\`  |               反斜杠               |

### 多行字符串

Go语言中要定义一个多行字符串时，就必须使用`反引号`字符：

```go
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```

### 字符串的常用操作

|                        方法                         |                             介绍                             |
| :-------------------------------------------------: | :----------------------------------------------------------: |
|                      len(str)                       |                     求长度（返回number）                     |
|          +或fmt.Sprintf("%s%s",str1,str2)           |                   拼接字符串(返回新字符串)                   |
|               strings.Split(str," ,")               |                    分割(返回分割后的 [ ])                    |
|             strings.Contains(str, "一")             |                  判断是否包含(返回Boolean)                   |
| strings.HasPrefix(str,""),strings.HasSuffix(str,"") | 前缀/后缀判断(返回Boolean，判断后面的字符是否是str字符串的前后缀) |
|         strings.Index(),strings.LastIndex()         |                        子串出现的位置                        |
|         strings.Join(a[]string, sep string)         |                           join操作                           |

```

	s1 := "字符串"
	str := "我是一个字符串！"
	
	length := len(s1)//求长度
	
	str1 :=s1+str//拼接
	
	str2 := fmt.Sprintf("%s+%s \n", s1, str) //?Sprintf是拼接后返回拼接的值
	
    fmt.Printf("str2: %s \n", str2)
		
	fmt.Printf("%s%s \n", s1, str)			//? Printf 是打印
	
	fmt.Println(strings.Split(str, ""))     //?字符串分割
	
	fmt.Println(strings.Contains(str, "一")) //? 字符串是否包含，返回Boolean
	
	str4 := "字符.txt"
	
	fmt.Println(strings.HasPrefix(str4, "字符"))//前缀
	
	fmt.Println(strings.HasSuffix(str4, "字符"))//后缀
	
	str5 := "abcdfc"
	
	fmt.Println(strings.Index(str5, "c"))     // ?第一次出现的位置
	
	fmt.Println(strings.LastIndex(str5, "c")) //? 最后一次出现的位置
	
	str6 := strings.Split(str, "")
	
	str7 := strings.Join(str6, "+")
	
	fmt.Println(str7)
	
	//str2: 字符串+我是一个字符串！
    //字符串我是一个字符串！
    //[我 是 一 个 字 符 串 ！]
    //true
    //true
    //false
    //2
    //5
    //我+是+一+个+字+符+串+！
```

## byte和rune类型

组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：

```go
var a := '中'
var b := 'x'
```

Go 语言的字符有以下两种：

1. **`uint8`类型，或者叫 byte 型，代表了`ASCII码`的一个字符。**
2. **`rune`类型，代表一个 `UTF-8字符`。**

当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`。

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。

```go
// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
```

输出：

```bash
104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³) 
104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河
```

**因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。**

**字符串底层是一个byte数组，所以可以和`[]byte`类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度**。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

## 修改字符串

**要修改字符串，需要先将其转换成`[]rune`（可以转多种文字）或`[]byte`（只能转英文），完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。**

### 因为字符串类型的变量的，存储内存地址是有部分公用的，一般来书字符串是不允许修改的，因为修改一个可能会影响其他的字符串，所以需要给他重新分配一个单独的内存的切片，再去修改，之后再转换成字符串

```go
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1) //['b','i','g',]
	byteS1[0] = 'p' //注意，这里是单引号的字符
	fmt.Println(string(byteS1)) //? pig

	s2 := "白萝卜"
	runeS2 := []rune(s2) //['白','萝','卜',]
	runeS2[0] = '红'//注意，这里是单引号的字符
	fmt.Println(string(runeS2))//?	红萝卜
    
    c1:='红'	//int32
    c2:="红" // string
}
```

## 类型转换

**Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。**

强制类型转换的基本语法如下：

```bash
T(表达式)
```

**其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等**.

比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。

```go
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c) //? 5
}
```

# if条件判断基本写法

`if`条件判断的格式如下：

```go
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

### if条件判断特殊写法

**if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，举个例子：**

```go
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
```

# for(循环结构)

for循环的基本格式如下：

```bash
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

条件表达式返回`true`时循环体不停地进行循环，直到条件表达式返回`false`时自动退出循环。

```go
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句可以被忽略，但是初始语句后的分号必须要写，例如：

```go
func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句和结束语句都可以省略，例如：

```go
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
```

### 无限循环

```go
for {
    循环体语句
}
```

### for循环可以通过`break`、`goto`、`return`、`panic`语句强制退出循环。

## goto(跳转到指定标签)

`goto`语句通过标签进行代码间的无条件跳转。`goto`语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用`goto`语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时：

```go
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}
```

**使用`goto`语句能简化代码：**

```go
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
    breakTag:
        fmt.Println("结束for循环")
}
```

## break(跳出循环)

**`break`语句可以结束`for`、`switch`和`select`的代码块。**

**`break`语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的`for`、`switch`和 `select`的代码块上。** 举个例子：

```go
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
```

## continue(继续下次循环)

**`continue`语句可以结束当前循环，开始下一次的循环迭代过程，仅限在`for`循环内使用。**

**在 `continue`语句后添加标签时，表示开始标签对应的循环**。例如：

```go
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
```

# for range(键值循环)

Go语言中可以使用`for range`遍历数组、切片、字符串、map 及通道（channel）。 通过`for range`遍历的返回值有以下规律：

1. 数组、切片、字符串**返回索引和值。**
2. map**返回键和值。**
3. 通道（channel）**只返回通道内的值**

```
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
```

# switch case

使用`switch`语句可方便地对大量的值进行条件判断。

```go
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
```

Go语言规定每个`switch`只能有一个`default`分支。

**一个分支可以有多个值，多个case值中间使用英文逗号分隔。**

```go
func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}
```

**分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：**

```go
func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
```

**`fallthrough`语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。**

```go
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough	//意思就是如果a通过，那b也会执行
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
```

输出：

```bash
a
b
```





# Array(数组)

数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。 基本语法：

```go
// 定义一个长度为3元素类型为int的数组a
var a [3]int
```

## 数组定义：

```bash
var 数组变量名 [元素数量]T
```

比如：`var a [5]int`， 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 `[5]int`和`[10]int`是不同的类型。

```go
var a [3]int
var b [4]int
a = b //不可以这样做，因为此时a和b是不同的类型
```

数组可以通过下标进行访问，下标是从`0`开始，最后一个元素下标是：`len-1`，访问越界（下标在合法范围之外），则触发访问越界，会panic。

## 数组的初始化

数组的初始化也有很多方式。

### 方法一

初始化数组时可以使用初始化列表来设置数组元素的值。

```go
func main() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}
```

### 方法二

按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：

```go
func main() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}
```

### 方法三

我们还可以使用指定索引值的方式来初始化数组，例如:

```go
func main() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}
```

## 数组的遍历

遍历数组a有以下两种方法：

```go
func main() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
```

## 多维数组

Go语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）。

### 二维数组的定义

```go
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}
```

### 二维数组的遍历

```go
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
```

输出：

```bash
北京	上海	
广州	深圳	
成都	重庆	
```

**注意：** 多维数组**只有第一层**可以使用`...`来让编译器推导数组长度。例如：

```go
//支持的写法
a := [...][2]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
//不支持多维数组的内层使用...
b := [3][...]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
```

## 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```go
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
func main() {
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}
```

**注意：**

1. 数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
2. `[n]*T`表示指针数组，`*[n]T`表示数组指针 。



# 切片

## 切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。



声明切片类型的基本语法如下：

```go
var name []T
```
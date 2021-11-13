package main

import (
	"fmt"
)

func NewFriend(sex int, age int, hobby string) (string, error){
	return "", nil
}


//////////////////////////////////////////////////////////////////


type Sex int
type Age int


func NewFriend2(hobby string, args ...interface{}) (string, error) {
	for _, arg := range args {
		switch arg.(type) {
		case Sex:
			fmt.Println(arg, "is sex")
		case Age:
			fmt.Println(arg, "is age")
		default:
			fmt.Println("未知的类型")
		}
	}
	return "", nil
}


//////////////////////////////////////////////////////////////////
// 选项设计模式


type Option func(*option)

// option 参数配置项
type option struct {
	sex int
	age int
}


func NewFriend3(hobby string, opts ...Option) (string, error) {
	opt := new(option)
	for _, f := range opts {
		f(opt)
	}

	fmt.Println(opt.sex, "is sex")
	fmt.Println(opt.age, "is age")

	return "", nil
}

func WithSex(sex int) Option {
	return func(opt *option) {
		opt.sex = sex
	}
}

func WithAge(age int) Option {
	return func(opt *option) {
		opt.age = age
	}
}

func main() {

	//friends, err := NewFriend3(
	//	"看书",
	//	WithAge(30),
	//	//WithSex(1),
	//)
	//
	//if err != nil {
	//	fmt.Println(friends)
	//}

}
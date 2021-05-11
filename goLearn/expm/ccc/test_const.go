package const_test

import "fmt"

const (
	BEIJIN = 10 * iota
	SHANGHAI
	SHENZHEN
)

func Test() {
	fmt.Println("Test....")
}

func init() {
	fmt.Println("lib1. init() ...")
}

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("1.txt", os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte("黄河之水天上来，奔流到海不复回.。\r\n"))
	}
}
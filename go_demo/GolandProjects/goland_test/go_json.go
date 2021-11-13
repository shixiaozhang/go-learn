package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	info := []Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}

	// 创建文件
	filePtr, err := os.Create("info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)

	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())

	} else {
		fmt.Println("编码成功")
	}

	// 打开文件
	filePtr2, err2 := os.Open("./info.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err2.Error())
		return
	}
	defer filePtr2.Close()
	var info2 []Website
	// 创建json解码器
	decoder := json.NewDecoder(filePtr2)
	err = decoder.Decode(&info2)
	if err2 != nil {
		fmt.Println("解码失败", err2.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info2)
	}
}

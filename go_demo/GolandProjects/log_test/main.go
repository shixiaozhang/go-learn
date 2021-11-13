package main

import (
	"log"
	"os"
)

// 配置log
func init() {
	// 设置前缀
	log.SetPrefix("JS: ")

	// 日志存文件
	// 创建|存在则附加|只能写  0666权限
	f, err := os.OpenFile("./js.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)
	// 设置格式 日期|时间|微妙|文件名
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// 记录日志
	log.Println("1234")
	// 记录日志并暂停程序
	log.Fatalln("1234")

	// 记录日志并抛出异常
	//log.Panicln("1234")

}

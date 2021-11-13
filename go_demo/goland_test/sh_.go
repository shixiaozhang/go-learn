package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

// 执行shell命令
func main() {
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("sh")
	cmd.Stdin = in
	go func() {
		in.WriteString("echo hello world > test.txt\n")
		in.WriteString("exit\n")
		
	}()
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
// 执行shell脚本并获取输出
func shell() {
	// 返回一个 cmd 对象
	cmd := exec.Command("sh", "-c", "./sh.sh")
	// 收返回值[]byte, error
	b,_:= cmd.Output()
	content:=string(b)
	fmt.Println(content)
	reg := regexp.MustCompile(`[0-9]+.[0-9]+ id`)
	fmt.Println(reg.FindAllString(content, -1)[0])
	shhh()
}

// 执行shell命令并获取输出
func shhh() {
	command := "top -b -n 1 | head"
	cmd := exec.Command("/bin/bash", "-c", command)
	fmt.Println(cmd)
	bytes,err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	resp := string(bytes)
	log.Println(resp)
}
// echo export NODE_ID=NODE_ID >> /etc/profile

// 设置命令超时
func CpuPercentSh() (float64, error) {
	command := "top -b -n 1 | head"
	cmd, err := Command("/bin/bash", "-c", command)
	if err != nil {
		return 0, err
	}
	// 收返回值[]byte, error
	content := string(cmd)
	reg := regexp.MustCompile(`([0-9]+.[0-9]+) id`)
	a := reg.FindStringSubmatch(content)[1]
	v2, _ := strconv.ParseFloat(a, 64)
	return 100.0 - v2, nil
}

func Command(name string, arg ...string) ([]byte, error) {
	ctxt, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	cmd := exec.CommandContext(ctxt, name, arg...)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Start(); err != nil {
		return buf.Bytes(), err
	}
	if err := cmd.Wait(); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

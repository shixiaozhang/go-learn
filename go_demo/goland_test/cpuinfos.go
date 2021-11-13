package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {

	cpuInfo, _ := cpu.Info()
	fmt.Println("cpu型号：",cpuInfo[0].ModelName)
	fmt.Println("cpu核心数：",cpuInfo[0].Cores)
	fmt.Println(strconv.FormatInt(int64(cpuInfo[0].Cores),10))

	hostinfo, _ := host.Info()
	fmt.Println("设备名称：",hostinfo.Hostname)
	format := time.Unix(int64(hostinfo.BootTime+hostinfo.Uptime), 0).Format("2006-01-02 15:04:05")
	fmt.Println("系统时间：",format)


	// 监测cpu占用率
	forever:=make(chan []float64,2)

	go func() {
		cpuPercent, _ := cpu.Percent(time.Second*2, true)
		forever<-cpuPercent
		//fmt.Println(cpuPercent)
	}()

	go func() {
		cpuPercent2, _ := cpu.Percent(time.Second*2, false)
		forever<-cpuPercent2
	}()
	for i := 0; i < 2; i++ {
		cpurate:=<-forever
		if len(cpurate)>1{
			fmt.Println("CPU各核占用率：",cpurate)

		}else {
			fmt.Println("CPU总占用率：",cpurate)
			float32s2 := strconv.FormatFloat(cpurate[0], 'E', -1, 64)
			fmt.Printf("%T",float32s2)
			fmt.Println(float32s2)
		}
	}

	ip,_:=GetLocalIP()
	fmt.Println("系统网络ip：",ip)

// wsl
	// 监测网络连通性
	addresses:=[]string{
		"127.0.0.1:81",
		"www.baidu.com:http",
		"192.168.10.128:3306",
		"gitlab.dhms.net:https"}
	PingTest(addresses)
	//address:="www.baidu.com:http"
	Mac()


}

// 获取本机ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// 网络连通性测试
func PingTest(addresses []string){
	var wg sync.WaitGroup

	for _,address:=range addresses{
		wg.Add(1)
		go func(add string) {
			defer wg.Done()
			d := net.Dialer{Timeout: time.Second*5}
			start:=time.Now()
			conn,err:=d.Dial("tcp",add)
			elapsed:=time.Since(start)
			if err!=nil{
				fmt.Println(add,"无法连接","耗时：",elapsed)
			}else{
				fmt.Println(add,"可以连接","耗时：",elapsed)
				conn.Close()
			}
		}(address)
	}
	wg.Wait()

	//return []string{}
}

func Mac(){
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	//for _, inter := range interfaces {
	//fmt.Println(inter.Name)
	inter := interfaces[0]
	mac := inter.HardwareAddr.String() //获取本机MAC地址
	fmt.Println("MAC = ", mac)
	fmt.Printf("%T",mac)
	err = os.Setenv("NODE_ID", mac)
	if err != nil {
		fmt.Println("ERROR:NODE_ID SET----", err.Error())
	}

	//}
}

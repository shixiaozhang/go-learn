package main

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
)


/*
influxdb虽然很多时候都可以通过SQL语句操作，但数据更新例外。
influxdb中数据更新只能重新Insert，且需要tags和时间戳相同，所以不建议大量更新数据。

极少出现删除数据的情况，删除数据基本都是清理过期数据。
 */


func main() {
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "YvLOjJyzwkvCzB89dME1AVxqAMXS4twd1n-r-Ichy16MjGeK8Q3LUYGtxKTXgFdwFC85PZJHygocsIdBHo_s2g=="
	const bucket = "bucket"
	const org = "organization"

	// 初始化influxdb客户端
	client := influxdb2.NewClient("http://192.168.16.134:8086", token)
	// 在最后关闭客户端
	defer client.Close()

	// 使用influxdb 协议写数据
	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)

	// write line protocol

	//writeAPI.WriteRecord(fmt.Sprintf("statt,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	//writeAPI.WriteRecord(fmt.Sprintf("statt,unit=temperature avg=%f,max=%f", 22.5, 45.0))
	//刷新写入
	//writeAPI.Flush()

	// create point using fluent style
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 26.2).
		AddField("max", 59.0).
		SetTime(time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()

	// 执行flux查询
	//query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h) |> filter(fn: (r) => r._measurement == \"statt\")", bucket)
	//// 获取查询客户端
	//queryAPI := client.QueryAPI(org)
	//// get QueryTableResult
	//result, err := queryAPI.Query(context.Background(), query)
	//if err == nil {
	//	// 迭代查询响应
	//	for result.Next() {
	//		// 注意何时更改组键
	//		if result.TableChanged() {
	//			fmt.Printf("table: %s\n", result.TableMetadata().String())
	//		}
	//		// 访问数据
	//		fmt.Printf("value: %v\n", result.Record().Value())
	//	}
	//	// 检查错误
	//	if result.Err() != nil {
	//		fmt.Printf("query parsing error: %\n", result.Err().Error())
	//	}
	//} else {
	//	panic(err)
	//}
}
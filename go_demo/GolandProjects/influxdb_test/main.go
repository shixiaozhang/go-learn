package main

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	// 使用InfluxDB服务器基本URL和身份验证令牌创建新客户端
	client := influxdb2.NewClient("http://192.168.16.134:8086", "YvLOjJyzwkvCzB89dME1AVxqAMXS4twd1n-r-Ichy16MjGeK8Q3LUYGtxKTXgFdwFC85PZJHygocsIdBHo_s2g==")
	// 使用阻塞写入客户机写入所需的存储桶
	writeAPI := client.WriteAPIBlocking("bucket", "organization")
	// 使用完全参数构造函数创建点
	//p := influxdb2.NewPoint("stat",
	//	map[string]string{"unit": "temperature"},
	//	map[string]interface{}{"avg": 24.5, "max": 45.0},
	//	time.Now())
	//// 立即写入数据
	//writeAPI.WritePoint(context.Background(), p)
	// 使用fluent样式创建数据
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45.0).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)

	// Or write directly line protocol
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	writeAPI.WriteRecord(context.Background(), line)

	// Get query client
	queryAPI := client.QueryAPI("my-org")
	// Get parser flux query result
	result, err := queryAPI.Query(context.Background(), `from(bucket:"my-bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			// Observe when there is new grouping key producing new table
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// read result
			fmt.Printf("row: %s\n", result.Record().String())
		}
		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	}
	// Ensures background processes finishes
	client.Close()
}

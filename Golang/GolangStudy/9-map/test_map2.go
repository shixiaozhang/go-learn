package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap 是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap, "China")

	//修改
	cityMap["USA"] = "DC"
	ChangeValue(cityMap)
	//遍历
	printMap(cityMap)

	fmt.Println("-------")
	// 集合操作（python中的字典）
	//country:=make(map[string]string,10)
	country := map[string]string{
		"one": "nanjing",
		"two": "shanghai",
	}
	country["Japan"] = "东京"
	// 删除
	delete(country, "Japan")
	// 判断元素是否在集合中
	capital, ok := country["American"]
	if ok {
		fmt.Println("American的首都是", capital)
	} else {
		fmt.Println("American首都不存在")
	}

}

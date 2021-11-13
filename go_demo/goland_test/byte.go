package main

import "fmt"

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func Add(keys ...string) {
	for k, key := range keys {
		fmt.Println(k,key)
	}
}

func main() {
	fmt.Println([]byte("qwee"))
	Add("asd")
}
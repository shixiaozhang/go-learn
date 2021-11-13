package main

import "fmt"

func main() {
	if a := 2; a != 3&&a!=1 {
		fmt.Println("!=3")
	} else {
		fmt.Println("=3")
	}
}

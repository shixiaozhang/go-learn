package main

import (
	"errors"
	"fmt"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}else {
		return f,nil
	}
}

func main() {
	result, err:= Sqrt(-1)
	fmt.Println(result)

	if err != nil {
		fmt.Println(err.Error())
	}

}
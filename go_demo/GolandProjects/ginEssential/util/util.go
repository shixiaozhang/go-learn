package util

import (
	"math/rand"
	"time"
)

// 生成随机数
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopASDFGHJKLQWERTYUIOP")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)

}

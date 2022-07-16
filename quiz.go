package main

import (
	"encoding/base64"
	"fmt"
)

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func main() {

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	whatIsIt, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}

	myString := string(whatIsIt)
	fmt.Printf("Decoded text: %s\n", reverseString(myString))
}

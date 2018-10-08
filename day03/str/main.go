package main

import (
	"fmt"
)

func main() {
	testString()
}

func testString() {
	var str = "hello"
	fmt.Printf("str[0]=%c\n", str[0])

	for i, item := range str {
		fmt.Printf("%d : %c\n", i, item)
	}

	var bs []byte = []byte(str)
	bs[0] = '0'
	str = string(bs)
	fmt.Println("after modify:", str)
	fmt.Printf("len(str)=%d\n", len(str))

	str = "kim孙志good"
	fmt.Printf("len(str)=%d\n", len(str))
	for i, item := range str {
		fmt.Printf("%d : %c\n", i, item)
	}

	var b rune = '中'
	fmt.Printf("b=%c\n", b)

	var rs []rune
	rs = []rune(str)
	fmt.Printf("str len:%d\n", len(rs)) 
	fmt.Printf("str :%s\n", rs)
	fmt.Printf("string(str) :%s\n", string(rs))
	//why len(str) = 13, but len(rs) = 9
	//str is string rs is []rune
	//utf code, a char have 1 or more byte, english is 1 byte
	//but one chinese is 3 byte 
}

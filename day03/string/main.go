package main

import "fmt"

func main() {
	//testString()
	fmt.Println("-------------\n", reverse("hello"))
	//fmt.Println("-------------\n", reverse("hello 中国")) // error, why
	fmt.Println("-------------\n", reverseChinese("hello 中国"))
	//var str = "hloolh"
	var str = "上海海上"
	fmt.Println(str, " is huiwen : ", huiwen(str))
}

func testString() {
	var str = "hello"
	fmt.Printf("str[0] = %c, len=%d\n", str[0], len(str))

	for index, val := range str {
		fmt.Printf("%d -- %c\n", index, val)
	}

	//str[0] = 'Y' // string is readonly
	//fmt.Println(str)
	//if you want to modify string, you should to convert to slice
	var byteSlice []byte
	byteSlice = []byte(str)
	fmt.Println(byteSlice)
	byteSlice[0] = 'Y'
	fmt.Println(byteSlice)

	str = string(byteSlice) // convert []byte to string
	fmt.Println(str)

	fmt.Printf("len(str) = %d\n", len(str))

	//一个中文字符占三个字节 : why
	//字符串是用字节数组来存储的
	str = "hello, 傻逼"
	fmt.Printf("len(str) = %d\n", len(str))

	var b rune = '中'
	fmt.Printf("b=%c\n", b)
	var runeSlice []rune
	runeSlice = []rune(str)
	fmt.Printf("str length = %d, len = %d\n", len(runeSlice), len(str))
}

func reverse(str string) string {
	var bytes []byte = []byte(str)
	for i := 0; i < len(str)/2; i++ {
		tmp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	//fmt.Println("inner = ", bytes)
	var rs = string(bytes)
	fmt.Println("inner = ", rs)
	return rs
}

func reverseChinese(str string) string {
	var r []rune = []rune(str)
	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	//fmt.Println("inner = ", bytes)
	//var rs = string(r)
	//fmt.Println("inner = ", rs)
	return string(r)
}

func huiwen(s string) bool {
	var str []rune = []rune(s) // why must convert to []rune
	n := len(str)
	result := true
	for i := 0; i < n/2; i++ {
		if str[i] == str[n-i-1] {
			continue
		} else {
			result = false
			break
		}
	}

	return result
}

package main

import (
	"fmt"
	"strings"
)

const (
	c1 int = iota
	c2
	c3
	c4 string = "c1-v"
	c5 string = "c5-v"
	c6 int    = iota
	c7 int    = 77
	c8
)

const (
	A = iota
	B
	C
	D = B
	E
	F = iota
	G
)

const (
	A1 = iota
	A2
)

func main() {
	var a int // default 0
	var b string
	var c bool
	var d int = 8
	var e string = "hello"
	var f float32

	fmt.Println(a, b, c, d, e, f)
	fmt.Printf("a=%d, b=%s, c=%t, d=%d, e=%s, f=%f\n", a, b, c, d, e, f)

	var (
		a1  int
		b1  bool
		c11 string
	)
	a1 = 122
	c11 = "asdfa"
	fmt.Println(a1, b1, c11)

	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8)

	const (
		a11 = 11 // you can skip the type
		s11 = "ericsson"
	)
	fmt.Println(a11, s11)
	fmt.Println(A, B, C, D, E, F, G)
	fmt.Println(A1, A2)

	var aa bool
	var bb bool = true
	if aa == true && bb == true {
		fmt.Println("both true")
	} else {
		fmt.Println("no true")
	}

	if aa == true || bb == true {
		fmt.Println("|| true")
	} else {
		fmt.Println("both false")
	}

	fmt.Printf("aa = %t, bb = %t\n", aa, bb)

	testInt()

	testString()
}

func testInt() {
	var a int8
	a = 18
	fmt.Println("a = ", a)
	a = -12
	fmt.Println("a = ", a)

	//a = 1243  // overflow
	//fmt.Println(a)

	//var b int
	var c float32 = 5.6
	fmt.Printf("c = %f\n", c)
}

func testString() {
	var a string = "hello world"
	a = a + a
	fmt.Println(a)
	fmt.Println("------")
	a = fmt.Sprintf("%s%s", a, a)
	fmt.Println(a)

	ips := "one:two:three"
	for _, ip := range strings.Split(ips, ":") {
		fmt.Println(ip)
	}

	arr := strings.Split(ips, ":")
	arrStr := strings.Join(arr, "##")
	fmt.Println(arrStr)

	fmt.Println(strings.HasPrefix(arrStr, "one"))
	fmt.Println(strings.HasSuffix(arrStr, "three"))
	fmt.Println(strings.Contains(arrStr, "one"))

	fmt.Printf("arrStr=%s\n", arrStr)
	fmt.Println("len = ", len(arrStr))

	fmt.Println(strings.Index(arrStr, "two"))

	arrStr = fmt.Sprintf("%s:two", arrStr)
	fmt.Printf("arrStr=%s\n", arrStr)
	fmt.Println("len = ", len(arrStr))
	fmt.Println(strings.LastIndex(arrStr, "two"))
}

//data-type
/*
1) bool
2) data type
	a) integer : byte
				 int, int8, int16, int32, int64,
				 rune
				 uint, uint8, uint16, uint32, uint64
	b) float : float32, float64
3) string
*/

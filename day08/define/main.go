package main

import (
	"fmt"
)

func main() {
	//var a [5]int // a is an int type array
	var a []int // a is a int type slice
	if a == nil {
		fmt.Println("a is nil")
	}
	//a[0] = 10 // will panic, because a is nil
	fmt.Println(a)
	oper()
	fmt.Println("-----------")
	sliceOper()
	fmt.Println("-----------")
	testSlice()
	fmt.Println("-----------")
	testSlice2()
}

//slice is a reference type
func oper() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:4] // slice b reference a from index 1 to index 3
	fmt.Println(b)
	b[1] = 100 // such modify will impact a because b reference to a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("-------")
	c := []int{6, 7, 8}//go will create a array and slice c will reference the array
	fmt.Println(c)
}

func sliceOper() {
	a := [5]int{1,2,3,4,5}
	var b []int
	b = a[1:4]
	fmt.Printf("slice b:%v\n", b)
	c := a[1:]
	fmt.Printf("slice c:%v\n", c)
	d := a[:3]
	fmt.Printf("slcie d:%v\n", d)
	e := a[:]
	fmt.Printf("slcie e:%v\n", e)

	//below address is the same, because b[0] and c[0] refer the same element of the array
	fmt.Printf("address b0: %v\n", &b[0])
	fmt.Printf("address c0: %v\n", &c[0])
}

func testSlice() {
	a := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("a : %v, type : %T\n", a, a)
	b := a[2:5]
	fmt.Printf("b : %v, type : %T\n", b, b)
	//b[0] += 10
	//b[1] += 10
	//b[2] += 10
	for index,_ := range b {
		b[index] += 10
	}
	fmt.Printf("a : %v, type : %T\n", a, a)
}

func testSlice2() {
	a := [...]int{1,2,3,4,5,6,7,8}
	s1 := a[:]
	s2 := a[:]
	s1[0] = 100
	s2[1] = 200
	fmt.Println(a)
	fmt.Printf("a-addr=%d, s1-addr=%d, s2-addr=%d\n",&a[1], &s1[1], &s2[1])
}
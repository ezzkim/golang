package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var f = Adder()
	fmt.Printf("f(1) = %d\n", f(1))
	fmt.Printf("f(3) = %d\n", f(3))
	fmt.Printf("f(5) = %d\n", f(4))

	/*
		//declare fixed length and init
		var a1 [3]int = [3]int{1, 2, 3}
		fmt.Println(a1)

		//declare fixed langth and init-portion
		a2 := [8]int{1, 2, 3, 4, 5}
		fmt.Println(a2)

		//declare no-fixed length and init, after init the length fixed
		a3 := [...]int{1, 2, 3, 4, 5, 6, 7}
		fmt.Println(a3)

		//declare fix-length and assign-init
		a4 := [5]int{3: 300, 4: 400}
		fmt.Println(a4)

		////array length is part of array type
		a5 := a4
		fmt.Println(a5, "length = ", len(a5))
	*/
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	//var a1 [3]int = [3]int{1, 2, 3}
	s := Sum(a[:])
	fmt.Printf("s = %d\n", s)

	testClosure()

	testSuffix()

	testCale()
	testThread()
	testThread2()
}

func Sum(a []int) int {
	var s int = 0
	var f = Adder()
	for _, item := range a {
		s = f(item)
	}

	return s
}

func Adder() func(int) int {
	var x int // the x init default value is 0
	return func(d int) int {
		x += d
		return x
	}
}

//////////////////////////////////start///////////////////
func baseAdd(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func testClosure() {
	f1 := baseAdd(10)
	fmt.Println(f1(1), f1(2))
	f2 := baseAdd(100)
	fmt.Println(f2(1), f2(2))
}
//////////////////////end//////////////////////

////////////////////////start//////////////////
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func testSuffix() {
	f1 := makeSuffixFunc(".bmp")
	f2 := makeSuffixFunc(".jpg")
	fmt.Println(f1("test"))
	fmt.Println(f2("test"))
}
///////////////////end///////////

/////////////start//////////////
func calc(base int) (func(int)int, func(int)int) {
	add := func(i int) int{
		base +=i 
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func testCale() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))
}
////////////////end/////////////

/////////////start//////////
func testThread() {
	for i:=0; i<5; i++ {
		go func() {
			fmt.Printf("i = %d\n", i)
		}()
	}

	time.Sleep(time.Second * 3)
}

func testThread2() {
	for i:=0; i<5; i++ {
		go func(n int) {
			fmt.Printf("n = %d\n", n)
		}(i)
	}

	time.Sleep(time.Second * 3)
}
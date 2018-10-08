package main

import (
	"fmt"
)

func main() {
	testDefer()
	testDefer2()
	testDefer3()
}

func testDefer() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("end")
}

func testDefer2() {
	for i:=0; i<5; i++ {
		defer fmt.Printf("i = %d\n", i)
	}

	fmt.Println("defer2 end")
}

func testDefer3() {
	var i int = 0
	defer fmt.Printf("defer2 i = %d\n", i)
	i = 100
	fmt.Printf("3 i = %d\n", i)

	fmt.Println("defer3 end")
}
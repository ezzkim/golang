package main

import (
	"fmt"
)

var ga int = 100

func main() {
	testG()
	testLoal()
}

func testG() {
	fmt.Printf("global ga = %d, %v\n", ga, &ga)
}

func testLoal() {
	var ga int = 200
	fmt.Printf("function ga = %d, %v\n", ga, &ga)
	{
		var ga int = 300
		fmt.Printf("block ga = %d, %v\n", ga, &ga)
	}
	fmt.Printf("function ga = %d, %v\n", ga, &ga)
}

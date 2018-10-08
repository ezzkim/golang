package main

import (
	"fmt"
	"time"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("end")
	time.Sleep(time.Second * 1)
	x, y := 1, 2
	fmt.Println(add(x, y))
}

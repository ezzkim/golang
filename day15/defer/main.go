package main

import (
	"fmt"
)

//return value  = x
// x + 1
//ret return value
func funcA() int {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

//x = 5
// x += 1
//ret
func funcB() (x int) {
	defer func() {
		x += 1
	}()
	return 5
}

//y = x
//x += 1
//ret command
func funcC() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

//x = 5
// copy x to innnter x and inner x += 1
//ret
func funcD() (x int) {
	defer func(x int) {
		x += 1
	}(x)
	return 5
}

func main() {
	fmt.Println(funcA())
	fmt.Println(funcB())
	fmt.Println(funcC())
	fmt.Println(funcD())
}

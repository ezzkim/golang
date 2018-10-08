package main

import (
	"fmt"
)

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	s := calc3(a...)
	fmt.Println(s)
	s = calc4(a[0], a[1:]...)
	fmt.Println(s)
	s = calc5(a[0], a[1], a[2:]...)
	fmt.Println(s)
	sm, sb := addSub(2, 3)
	fmt.Printf("%d : %d", sm, sb)

	s1, s2, s3, s4, s5 := 1, 2, 3, 4, 5
	ss1 := calc3(s1, s2, s3, s4, s5)
	ss2 := calc4(s1, s2, s3, s4, s5)
	ss3 := calc5(s1, s2, s3, s4, s5)
	fmt.Println("ss1 = ", ss1, "ss2 = ", ss2, "ss3 = ", ss3)
}

func addSub(a, b int) (int, int) {
	/*
		sum := a + b
		sub := a - b
		return sum, sub
	*/
	return a + b, a - b
}

//the vary number of parameter, at least one parameter
func calc3(b ...int) int {
	sum := 0
	for i, v := range b {
		fmt.Printf("%d -- %d\n", i, v)
		sum += v
	}

	return sum
}

//at lease transfer one parameter, at least two parameter
func calc4(a int, b ...int) int {
	sum := a
	for i, v := range b {
		fmt.Printf("%d -- %d\n", i, v)
		sum += v
	}

	return sum
}

//at lease transfer two parameter, at least three parameter
func calc5(a int, b int, c ...int) int {
	sum := (a + b)
	for i, v := range c {
		fmt.Printf("%d -- %d\n", i, v)
		sum += v
	}

	return sum
}

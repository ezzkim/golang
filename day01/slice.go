package main

import (
	"fmt"
)

func sliceTest() {
	s11 := []int{1, 2, 3, 4, 5}
	s12 := []int{7, 8, 9}
	fmt.Println("content = ", s11, ", len = ", len(s11), ", cap = ", cap(s11))
	copy(s12, s11)
	fmt.Println(s12)
	s12 = append(s12, 11, 22, 33, 44, 55)
	fmt.Println(s12)
	copy(s12[2:], s11[3:5])
	fmt.Println(s12)

	ss := make([]int, 10)
	//ss := make([]int, 0)
	ss = append(ss, s12...)
	fmt.Println("content = ", ss, ", len = ", len(ss), ", cap = ", cap(ss))
}

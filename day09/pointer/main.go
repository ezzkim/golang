package main

//value type : normal varity save the value
//pointer type : save the address of some value

import (
	"fmt"
)

func main() {
	var b int32
	b = 132
	var a *int32
	a = &b
	fmt.Printf("&b = %p\n", &b)
	fmt.Printf("&a = %p, a = %p, *a = %d\n", &a, a, *a)

	change(a)
	fmt.Printf("&a = %p, a = %p, *a = %d\n", &a, a, *a)

	aa := [...]int{1, 2, 3}
	fmt.Println(aa)
	modify_arr(&aa)
	fmt.Println(aa)

	testPointer()

	copy()
}

func change(v *int32) {
	*v = 555
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

//a is a pointer point a [3]int array
func modify_arr(a *[3]int) {
	a[1] = 100
}

//here a is slice, a reference type
func modify_arr2(a []int) {
	a[1] = 100
}

//make : slice, map, chan
//new : any type
func testPointer() {
	var a *int = new(int)
	*a = 100
	fmt.Printf("a = %d\n", *a)

	//b is slice pointer
	var b *[]int = new([]int) // b point a slice, but the sliceis nil now
	fmt.Printf("*b = %v\n", *b)
	if *b == nil {
		fmt.Printf("*b is nil now\n")
	}
	(*b) = make([]int, 5, 10) // if do not make, below will
	(*b)[0] = 1
	(*b)[1] = 2
	fmt.Printf("*b = %v\n", *b)
}

//value copy
//reference copy
func copy() {
	var a int = 23
	modify(a)
	fmt.Printf("a = %d\n", a)
	modify2(&a)
	fmt.Printf("a = %d\n", a)

}

func modify(a int) {
	a = 100
}

func modify2(a *int) {
	*a = 100
}

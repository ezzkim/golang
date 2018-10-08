package main

import (
	"fmt"
)

func main() {
	showPointer()
	var a int = 23
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)

	var b int = 321
	fmt.Printf("a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func showPointer() {
	var a int = 123
	fmt.Printf("a=%d, &a=%p\n", a, &a)
	var p *int = &a
	fmt.Printf("p=%p, *p=%d\n", p, *p)
}

func modify(p *int) {
	*p = 1000
}

func swap(a, b *int) {
	/*
	tmp := *a
	*a = *b
	*b = tmp
	*/
	*a, *b = *b, *a
}

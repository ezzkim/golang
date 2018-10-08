package main

import (
	"fmt"
)

func main() {
	testMake()
	fmt.Println("--------")
	testSlice()
	fmt.Println("--------")
	sliceExtense()
	fmt.Println("--------")
	sliceCopy()
}

func testMake() {
	var a []int
	//a = make([]int, 5, 10)
	a = make([]int, 5) // default the cap is the init length
	fmt.Printf("a=%v, len=%d, cap=%d, a-address = %v\n", a, len(a), cap(a), &a[0])
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	//a[4] = 50
	fmt.Printf("a=%v, len=%d, cap=%d, a-address = %v\n", a, len(a), cap(a), &a[0])
	//a[5] = 60 // panic
	a = append(a, 100) // the default enlarge the origin cap*2
	fmt.Printf("a=%v, len=%d, cap=%d, a-address = %v\n", a, len(a), cap(a), &a[0])
	a = append(a, 200)
	fmt.Printf("a=%v, len=%d, cap=%d, a-address = %v\n", a, len(a), cap(a), &a[0])
}

func testSlice() {
	a := [...]string{"a", "b", "c", "d", "e", "f"}
	b := a[1:3]
	fmt.Printf("b:%v, len:%d, cap:%d\n", b, len(b), cap(b))
	//the len is 2, but why the cap is 5, because the b slice is the reference of a
	b = append(b, "append1")
	//the slice b have attach the array, the append operation will impact the array
	fmt.Printf("b:%v, len:%d, cap:%d\n", b, len(b), cap(b))
	fmt.Println(a)

	fmt.Printf("a[1]-addr=%p, slice-addr=%p\n", &a[1], b)

	////////
	var es []int
	fmt.Printf("addr:%p, len=%d, cap=%d\n", es, len(es), cap(es))
	if es == nil {
		fmt.Printf("slice is empty\n")
	}
	//es[0] = 100 //will panic if es is nil

	fmt.Println("----------")

	//each time do append action, if the slice is nil or cap is full, will malloc the memory
	//the enlarge cap is the double of origin cap
	es = append(es, 100)
	fmt.Printf("addr:%p, len=%d, cap=%d\n", es, len(es), cap(es))
	es = append(es, 200)
	fmt.Printf("addr:%p, len=%d, cap=%d\n", es, len(es), cap(es))
	es = append(es, 300)
	fmt.Printf("addr:%p, len=%d, cap=%d\n", es, len(es), cap(es))
	es = append(es, 400)
	fmt.Printf("addr:%p, len=%d, cap=%d\n", es, len(es), cap(es))
	es = append(es, 500)
	fmt.Printf("addr:%p, len=%d, cap=%d\n", es, len(es), cap(es))
}

func sliceExtense() {
	var a []int = []int{1, 2, 3}
	var b []int = []int{4, 5, 6}

	a = append(a, 22, 33, 44)
	fmt.Printf("a = %v\n", a)
	a = append(a, b...) // entend the content of b
	fmt.Printf("a = %v\n", a)

	s := sliceParameter(a[:])
	fmt.Printf("sum = %d\n", s)

	fmt.Printf("a : %v\n", a)
	sliceModify(a[:]) // a will been modify becaue the slice is refernce
	fmt.Printf("a : %v\n", a)
}

func sliceParameter(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum += v
	}

	return sum
}

func sliceModify(a []int) {
	a[0] = 1000
}

func sliceCopy() {
	var a []int = []int{1, 2, 3}
	var b []int = []int{4, 5, 6}
	copy(a, b)
	fmt.Printf("a=%v, b=%v\n", a, b)

	var a1 []int = []int{1, 2}
	var b1 []int = []int{4, 5, 6}
	copy(a1, b1)
	fmt.Printf("a1=%v, b1=%v\n", a1, b1)

	var a2 []int = []int{1, 2, 3, 14, 15}
	var b2 []int = []int{4, 5, 6}
	copy(a2, b2)
	fmt.Printf("a2=%v, b2=%v\n", a2, b2)
}

//make only support : slice, map, chan
//new support 

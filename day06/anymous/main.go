package main

import (
	"fmt"
)

func main() {
	testFunc()
	testFunc2()
	testFunc3()
	testFunc4()
	testFunc5()
}

func add(a, b int) int {
	return a + b
}

func testFunc() {
	f1 := add
	fmt.Printf("type : %T\n", f1)
	fmt.Printf("type : %v\n", f1)
	fmt.Printf("type : %+v\n", f1)
	fmt.Println(f1(1, 5))
}

func testFunc2() {
	var i int = 0
	defer fmt.Printf("defer i = %d\n", i) // i = 0
	/*
		defer func() {

		}()
	*/
	i = 100
	fmt.Printf("func end, i=%d\n", i)
}

func testFunc3() {
	var i int = 0
	//defer fmt.Printf("defer i = %d\n", i)  // i = 0
	defer func() {
		fmt.Printf("defer i = %d\n", i) //100, clusure
	}()

	i = 100
	fmt.Printf("func end, i=%d\n", i)
}

func testFunc4() {
	var i int = 0
	//defer fmt.Printf("defer i = %d\n", i)  // i = 0
	defer func(a int) {
		fmt.Printf("defer i = %d\n", a) // 0
	}(i) // i value have been transfer to defer func

	i = 100
	fmt.Printf("func end, i=%d\n", i)
}

func testFunc5() {
	s1 := calc(100, 300, add)
	s2 := calc(100, 300, sub)
	fmt.Printf("s1 = %d, s2 = %d\n", s1, s2)
}

func calc(a, b int, op func(int, int) int) int {
	fmt.Printf("call function name : %T\n", op)
	return op(a, b)
}

/*
func add(a, b int) int {
	return a + b
}*/

func sub(a, b int) int {
	return a - b
}

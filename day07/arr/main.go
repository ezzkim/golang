package main

import (
	"fmt"
)

func main() {
	// define a int arry, the length is 5
	var a1 [5]int 
	a1[2] = 22
	fmt.Println(a1)

	// define a int arry, the length is 5, and the init value is 1,2,3,4,5
	var a2 [5]int = [5]int{1,2,3,4,5}
	fmt.Println(a2)

	//declare fixed langth and init-portion
	a3 := [8]int{1, 2, 3, 4, 5}
	fmt.Println(a3)

	//declare no-fixed length and init, after init the length fixed
	a4 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a4)

	//declare fix-length and assign-init
	a5 := [5]int{3: 300, 4: 400}
	fmt.Println(a5)

	//var b [6]int
	var b [5]int
	b = a5 // if the length of b is not 5, you can assign a5 to b
	fmt.Println(b)
}

//in go, the length is the part of array type
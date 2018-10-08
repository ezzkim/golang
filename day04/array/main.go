package main

import (
	"fmt"
)

func main() {
	testArray()
}

func testArray() {
	var a1 [3]int
	a1[0] = 11
	a1[2] = 33
	fmt.Println(a1)

	a2 := [8]int{1, 2, 3, 4, 5}
	fmt.Printf("a = %v, len = %d\n", a2, len(a2))

	//append(a2, 22) // must slice
	//fmt.Println(a2)

	a3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a3 = %v, len = %d\n", a3, len(a3))

	a4 := [5]int{3: 300, 4: 400} // a4[3] = 300, a4[4] = 400
	fmt.Println(a4)

	////array length is part of array type
	a5 := a4 // copy array
	fmt.Println(a5, "length = ", len(a5))

	travel(a5[:])

	var aa [3][2]int
	aa[0][0] = 10
	aa[0][1] = 20
	aa[1][0] = 11
	aa[1][1] = 21
	aa[2][0] = 12
	aa[2][1] = 22
	travel(aa[0][:])
	travel(aa[1][:])
	travel(aa[2][:])

	ac := [...]int{1, 3, 5, 8, 7}
	fmt.Println("ac = ", ac)
	s := 8
	ns := getSumIndex(ac[:], s)
	for k, v := range ns {
		fmt.Printf("ac[%d] + ac[%d] = %d\n", k, v, s)
	}

	for _, one := range getSumIndex2(ac[:], s) {
		fmt.Printf("%s\n", one)
	}
}

func travel(a []int) {
	for _, item := range a {
		fmt.Printf("%d ", item)
	}
	fmt.Printf("\n")
}

func getSumIndex(a []int, s int) map[int]int {
	n := len(a)
	ns := make(map[int]int)
	for i := 0; i <= n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i]+a[j] == s {
				ns[i] = j
			}
		}
	}

	return ns
}

func getSumIndex2(a []int, s int) []string {
	n := len(a)
	ns := make([]string, 0)
	for i := 0; i <= n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i]+a[j] == s {
				ns = append(ns, fmt.Sprintf("a[%d] + a[%d] = %d", i, j, s))
			}
		}
	}

	return ns
}

package main

import (
	"sort"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	showResult()
	testSort()
}

func showResult() {
	var sa = make([]string, 5, 10) // sa have 5 init empty string
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i)) // append just add data to the end of slice
	}
	fmt.Println(sa)
	for _, item := range sa {
		fmt.Printf("[%s] ", item)
	}
	fmt.Printf("\n")
}

func testSort() {
	a := [...]int{5,4,3,2,1}
	fmt.Printf("a = %v\n", a)
	sort.Ints(a[:])
	fmt.Printf("a = %v\n", a)

	rand.Seed(time.Now().Unix())
	var aa [10]int
	for i:=0; i<len(aa); i++ {
		aa[i] = rand.Intn(1000)
	}
	fmt.Printf("aa = %v\n", aa)
	sort.Ints(aa[:])
	fmt.Printf("aa = %v\n", aa)

	s := [...]string{"w", "z", "one"}
	fmt.Printf("s = %v\n", s)
	sort.Strings(s[:])
	fmt.Printf("s = %v\n", s)
}

func passGen(){
	
}

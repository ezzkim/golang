package main

import (
	"fmt"

	"github.com/study/golang/day06/calc"
)

func main() {
	var s1 int = 200
	var s2 int = 300

	sum := calc.Add(s1, s2)
	fmt.Printf("calc.A=%d, sum=%d\n", calc.A, sum)
	//fmt.Printf("calc.a=%d, sum=%d", calc.a, sum)
	fmt.Printf("calc.FetchA=%d, calc.TestSub=%d\n", calc.Fetcha(), calc.Testsub(s2,s1))
}

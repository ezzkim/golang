package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("sum = %d\n", sumArray(a[:]))

	ac := [...]int{1, 3, 5, 8, 7}
	fmt.Println("ac = ", ac)
	for _, item := range sumIndex(ac[:], 8) {
		fmt.Println(item)
	}

	randArray()
}

func sumArray(a []int) int {
	fmt.Printf("parameter type = %T\n", a)
	var sum int = 0
	for _, item := range a {
		sum += item
	}

	return sum
}

func sumIndex(a []int, v int) []string {
	strArray := make([]string, 0)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == v {
				one := fmt.Sprintf("index[%d] + index[%d] = %d", i, j, v)
				strArray = append(strArray, one)
			}
		}
	}

	return strArray
}

func randArray() {
	rand.Seed(time.Now().Unix())
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(1000)
	}
	fmt.Println(a)
}

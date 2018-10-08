package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := [...]int{1,2,3,45,5}
	showArray(a[:])

	var aa [3][2]int = [3][2]int {
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println(aa)
	showArray(aa[0][:])
	showArray(aa[1][:])
	showArray(aa[2][:])
	fmt.Printf("----------\n")
	mutil()
}

func showArray(a []int) {
	/*
	for i:=0; i<len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Printf("\n")
	*/
	for _, item := range a {
		fmt.Printf("%d ", item)
	}
	fmt.Printf("\n")
}

func mutil() {
	var aa [3][2]int
	aa[0][0] = 10
	aa[0][1] = 20
	aa[1][0] = 11
	aa[1][1] = 21
	aa[2][0] = 12
	aa[2][1] = 22
	showArray(aa[0][:])
	showArray(aa[1][:])
	showArray(aa[2][:])

	as := [3][2]string{
		{"one-1", "one-2"},
		{"two-1", "two-2"},
		{"three-1", "three-2"},
	}

	fmt.Println(as[2][1])
	copy()

	ia := [...]int{1, 2, 3, 4, 5}
	sum := getSum(ia[:])
	fmt.Println("sum = ", sum)

	ac := [...]int{1, 3, 5, 8, 7}
	fmt.Println("ac = ", ac)
	getSumIndex(ac[:], 8)
	fmt.Println("--------------")
	str := getSumIndex2(ac[:], 8)
	fmt.Println(str)
	for _, one := range getSumIndex3(ac[:], 8) {
		fmt.Println(one)
	}

	testRandSum()
}

//array copy is value copy not reference copy
func copy() {
	a := [3]int{10, 20, 30}
	b := a // array is value type, b copy a new array
	b[0] = 100
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}

func getSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	/*
		for i:=0;i<len(a); i++ {
			sum += a[i]
		}*/

	return sum
}

//func getSumIndex(a []int, sum int) [...][2]int {
func getSumIndex(a []int, sum int) {
	//var array [...][2]
	//return array
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == sum {
				fmt.Printf("sum(%d) = index(%d) + index(%d)\n", sum, i, j)
			}
		}
	}
}

func getSumIndex2(a []int, sum int) string {
	//var array [...][2]
	//return array
	var str string
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == sum {
				str += fmt.Sprintf("index(%d) + index(%d) = %d\n", i, j, sum)
			}
		}
	}

	return str
}

func getSumIndex3(a []int, sum int) []string {
	//var array [...][2]
	//return array
	str := make([]string, 0)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == sum {
				str = append(str, fmt.Sprintf("index(%d) + index(%d) = %d", i, j, sum))
			}
		}
	}

	return str
}

func testRandSum() {
	rand.Seed(time.Now().Unix())  //init an seed other each time the value is the same
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000) // between 0 to 999
		//rand.Int() // 0 to int-Max value
	}
	fmt.Println(b)
	sum := getSum(b[:])
	fmt.Println("sum = ", sum)
}

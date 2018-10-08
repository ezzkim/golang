package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testIf()
	testFor()
	testSwitch()
	showTable()
}

func testIf() {
	rand.Seed(time.Now().Unix())
	ri := rand.Intn(100)
	fmt.Printf("rand data = %d\n", ri)
	if ri%2 == 0 {
		fmt.Println("ou shu")
	} else {
		fmt.Println("ji shu")
	}
}

func testFor() {
	str := "ericsson"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d : %c\n", i, str[i])
	}

	for _, ch := range str {
		fmt.Printf("%c\n", ch)
	}
}

func testSwitch() {
	var i = 12
	switch i {
	case 1, 2, 3, 4:
		fmt.Println("1-4")
	case 10, 11, 12, 13:
		fmt.Println("10-13")
		fallthrough
	default:
		fmt.Println("default")
	}
}

func showTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d", j, i, i*j)
			if j != i {
				fmt.Print(",")
			}
		}
		fmt.Println("")
	}
}

package main

import (
	"fmt"
)

func main() {
	verify()
	shuixian()
	example3()
}

func justify(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func verify() {
	a := make([]int, 0)
	for i := 1; i < 101; i++ {
		if justify(i) {
			a = append(a, i)
		}
	}

	fmt.Println("1-100 prime : ", a)
}

func is_shuixian(n int) bool {
	b := n / 100
	s := n / 10 % 10
	g := n % 10

	sum := g*g*g + s*s*s + b*b*b

	return n == sum
}

func shuixian() {
	a := make([]int, 0)
	for i := 100; i < 1000; i++ {
		if is_shuixian(i) {
			a = append(a, i)
		}
	}

	fmt.Println("100-1000 shuxian : ", a)
}

func calc(str string) (charCount int, numCount int, spaceCount int, otherCount int) {

	utfChars := []rune(str)
	for i := 0; i < len(utfChars); i++ {
		if utfChars[i] >= 'a' && utfChars[i] <= 'z' || utfChars[i] >= 'A' && utfChars[i] <= 'Z' {
			charCount++
			continue
		}

		if utfChars[i] >= '0' && utfChars[i] <= '9' {
			numCount++
			continue
		}

		if utfChars[i] == ' ' {
			spaceCount++
			continue
		}

		otherCount++
	}
	return
}

func example3() {
	var str string = "dfka   我爱天安门 38333"
	charCount, numCount, spCount, other := calc(str)
	fmt.Printf("char count:%d num count:%d sp count:%d other:%d\n", charCount, numCount, spCount, other)
}

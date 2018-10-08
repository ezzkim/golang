package util

import (
	"fmt"
)

//exec first
func init() {
	fmt.Printf("util init 2\n")
}

//exec second
func init() {
	fmt.Printf("util init\n")
}

func swap(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func SelectSort(a []int) {
	for i:=0; i<len(a); i++ {
		index := i
		for j:=i+1; j<len(a); j++ {
			if a[index] > a[j] {
				index = j
			}
		}
		if index != i {
			swap(&a[i], &a[index])
		}
	}
}

func BubbleSort(a []int) {
	var bSwap bool
	var n int = len(a)
	for {
		bSwap = false
		for i:=0; i<n-1; i++ {
			if a[i] > a[i+1] {
				swap(&a[i], &a[i+1])
				bSwap = true
			}
		}
		if !bSwap {
			break
		} else {
			n--
		}
	}
}

func InsertSort(a []int) {
	var tmp int
	var pos int
	for i:=1; i<len(a); i++ {
		pos = i
		tmp = a[i]
		for j:=i-1; j>=0; j-- {
			if a[j] > tmp {
				a[j+1] = a[j] 
				pos--
			}
		}
		a[pos] = tmp
	}
}
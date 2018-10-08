package main

import (
	"fmt"
)

func main() {
	var str string = "afsa adf fds 美国 323 3"
	cn,dn,sn,on := calc(str)
	fmt.Println(cn,dn,sn,on)
}

func calc(str string) (int, int, int, int) {
	var cn int = 0
	var dn int = 0
	var sn int = 0
	var on int = 0

	utfChars := []rune(str)
	for _, item := range utfChars {
		if (item >= 'a' && item <= 'z') || (item >= 'A' && item <= 'Z') {
			cn++
		} else if item >= '0' && item <= '9' {
			dn++
		} else if item == ' ' {
			sn++
		} else {
			on++
		}
	}

	return cn, dn, sn, on
}

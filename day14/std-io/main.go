package main

import (
	"fmt"
	"os"
)

//os.Stdin
//os.Stdout
//os.Stderr

func main() {
	fmt.Println("hello, input a string") // Fprintln(os.Stdout, a...)
	var buf [16]byte
	os.Stdin.Read(buf[:])

	//fmt.Println(string(buf[:]))
	os.Stdout.WriteString(string(buf[:]))

	os.Stderr.WriteString(string(buf[:]))
}

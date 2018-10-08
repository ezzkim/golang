package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Fscanf(file, format, a...)
	var a int
	var b string
	var c float32
	//fmt.Fscanf(os.Stdin, "%d%s%f", &a,&b,&c)
	fmt.Fscan(os.Stdin, &a, &b, &c)

	//fmt.Println(a, b, c)
	//Fprintf
	fmt.Fprintln(os.Stdout, "stdout : ", a, b, c)
}

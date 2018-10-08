package main

import (
	"fmt"
)

//Scanf
//Scan
//Scanln

func main() {
	//testScanf()
	//testScan()
	//testScanln()
	//testSscanf()
	//testSscan()
	testSscanln()
}

func testScanf() {
	var a int
	var b string
	var c float32

	fmt.Printf("input int string float32\n")
	//fmt.Scanf("%d%s%f", &a, &b, &c)
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%s\n", &b)
	fmt.Scanf("%f\n", &c)
	fmt.Printf("a[%d], b[%s], c[%f]\n", a, b, c)
}

func testScan() {
	var a int
	var b string
	var c float32

	fmt.Printf("input int string float32\n")
	fmt.Scan(&a, &b, &c)
	//fmt.Scanf("%s\n", &b)
	//fmt.Scanf("%f\n", &c)
	fmt.Printf("a[%d], b[%s], c[%f]\n", a, b, c)
}

func testScanln() {
	var a int
	var b string
	var c float32

	fmt.Printf("input int string float32\n")
	fmt.Scanln(&a, &b, &c)
	//fmt.Scanf("%s\n", &b)
	//fmt.Scanf("%f\n", &c)
	fmt.Printf("a[%d], b[%s], c[%f]\n", a, b, c)
}

func testSscanf() {
	var a int
	var b string
	var c float32

	var str string = "123 one 123.43"
	//fmt.Printf("input int string float32\n")
	//fmt.Scanf("%d%s%f", &a, &b, &c)
	fmt.Sscanf(str, "%d%s%f", &a, &b, &c)
	fmt.Printf("a[%d], b[%s], c[%f]\n", a, b, c)
}

func testSscan() {
	var a int
	var b string
	var c float32

	var str string = "123 one\n\n 123.43"
	//fmt.Printf("input int string float32\n")
	//fmt.Scanf("%d%s%f", &a, &b, &c)
	fmt.Sscan(str, &a, &b, &c)
	fmt.Printf("a[%d], b[%s], c[%f]\n", a, b, c)
}

func testSscanln() {
	var a int
	var b string
	var c float32

	var str string = "123 one \n123.43\n"
	//fmt.Printf("input int string float32\n")
	//fmt.Scanf("%d%s%f", &a, &b, &c)
	fmt.Sscanln(str, &a, &b, &c)
	fmt.Printf("a[%d], b[%s], c[%f]\n", a, b, c)
}

//output to terminal
//Printf
//Println
//Print

//output to string
//Sprintf
//Sprintln
//Sprint

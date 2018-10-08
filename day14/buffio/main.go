package main

import (
	"bufio"
	"fmt"
	//"bufio"
	"os"
)

func main() {
	/*
	var str string
	//how to read one line
	fmt.Scanf("%s", &str) // bad methid
	fmt.Printf("oneline = [%s]\n", str)
	*/
	
	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	fmt.Printf("read from bufio:[%s]\n", str)
}

package main

import (
	"flag"
	"fmt"
	"os"
)

var recusive bool
var test string
var level int

func main() {
	//os.Args is simple method to get command line parameters
	fmt.Printf("args[0]=%s\n", os.Args[0])
	if len(os.Args) > 1 {
		for index, v := range os.Args {
			if index == 0 {
				continue
			}
			fmt.Printf("args[%d]=%v\n", index, v)
		}
	}

	fmt.Println("---------------")
	fmt.Printf("resusive = %v\n", recusive)
	fmt.Printf("test = %v\n", test)
	fmt.Printf("level = %v\n", level)
}

func init() {
	flag.BoolVar(&recusive, "r", false, "recusive method")
	flag.StringVar(&test, "t", "default-string", "string option")
	flag.IntVar(&level, "l", 1, "level of xxxx")

	flag.Parse()
}

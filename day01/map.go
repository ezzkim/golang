package main

import (
	"fmt"
)

func reverseMap() {
	mm1 := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	for k, v := range mm1 {
		fmt.Printf("%d : %s\n", k, v)
	}	
	fmt.Println("------------------")
	mm2 := make(map[string]int)
	for k, v := range mm1 {
		mm2[v] = k
	}
	for k, v := range mm2 {
		fmt.Printf("%s : %d\n", k, v)
	}
}

func mapTest() {
	fmt.Println("---begin map test---")

	reverseMap()

	fmt.Println("---end map test---")
}

package main

import (
	"fmt"
)

func main() {
	names := map[string]int{
		"Matthew":   0,
		"Sarah":     0,
		"Augustus":  0,
		"Heidi":     0,
		"Emile":     0,
		"Peter":     0,
		"Giana":     0,
		"Adriano":   0,
		"Aaron":     0,
		"Elizabeth": 0,
	}
	fmt.Println(names)
	var coins int = 50

	for name, _ := range names {
		var num int = 0
		for _, ch := range name {
			switch ch {
			case 'a', 'A':
				num += 1
			case 'e', 'E':
				num += 1
			case 'i', 'I':
				num += 2
			case 'o', 'O':
				num += 3
			case 'u', 'U':
				num += 5
			}
		}
		names[name] = num
		coins -= num
	}

	for name, num := range names {
		fmt.Printf("%s : %d\n", name, num)
	}

	fmt.Printf("Left coins number : %d\n", coins)
}

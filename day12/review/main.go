package main

import (
	"fmt"
)

var (
	names = [...]string{
		"Matthew",
		"Sarah",
		"Augustus",
		"Heidi",
		"Emile",
		"Peter",
		"Giana",
		"Adriano",
		"Aaron",
		"Elizabeth",
	}

	coins = 50
)

func main() {
	fmt.Printf("coins = %d, names : %v\n", coins, names)

	coinsMap := make(map[string]int, len(names))
	for _, name := range names {
		num := calcCoinsNumber(name)
		coinsMap[name] = num
		coins -= num
	}

	for k, v := range coinsMap {
		fmt.Printf("%s : %d\n", k, v)
	}
	fmt.Printf("left coins : %d\n", coins)
}

func calcCoinsNumber(name string) int {
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

	return num
}

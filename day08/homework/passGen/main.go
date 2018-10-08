package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "!@#$%^&*"
)

func main() {
	rand.Seed(time.Now().Unix())
	parseArgs()
	fmt.Printf("length=%d charset=%s\n", length, charset)

	//str := genCharSet()
	//fmt.Println(str)
	passwd := generatePassword()
	fmt.Printf("password : %s\n", passwd)
}

func parseArgs() {
	flag.IntVar(&length, "l", 10, "the length os password")
	flag.StringVar(&charset, "t", "num", "num[0-9] char[a-z A-Z], mix[0-9 a-z A-Z] advance[0-9 a-z A-Z special-chars]")
	flag.Parse()
}

func genCharSet() string {
	/*
		s := make([]byte, 0)
		for ch:='A'; ch<='Z'; ch++ {
			s = append(s, byte(ch))
		}
		for ch:='a'; ch<='z'; ch++ {
			s = append(s, byte(ch))
		}
	*/

	s := make([]rune, 0)
	for ch := 'A'; ch <= 'Z'; ch++ {
		s = append(s, ch)
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		s = append(s, ch)
	}

	return string(s)
}

func generatePassword() string {
	var passwd []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}

	fmt.Printf("source:%s\n", sourceStr)

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr)) // gen index
		passwd[i] = sourceStr[index]
	}

	return string(passwd)
}

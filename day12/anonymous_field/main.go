package main

import (
	"fmt"
)

type User struct {
	Username  string
	Sex       string
	int
	string
}

func main() {
	var user User
	user.Username = "user01"
	user.Sex = "man"
	user.int = 18
	user.string = "hello world"

	fmt.Printf("user=%#v\n", user)
}
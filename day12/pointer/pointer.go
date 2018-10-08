package main

import (
	"fmt"
)

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user *User
	fmt.Printf("user=%v\n", user)

	var user01 *User = &User{}  // 
	//user01.Age = 18
	//user01.Username = "user01"
	//the same as below
	(*user01).Age = 18
	(*user01).Username = "user01"


	var user02 *User = &User{
		Age: 18,
		Username: "user02",
	}

	fmt.Printf("user01:%p\n", user01)
	fmt.Printf("user01:%#v\n", user01)
	fmt.Printf("user02:%p\n", user02)
	fmt.Printf("user02:%#v\n", user02)

	var user03 *User = new(User)  //
	user03.Age = 18
	user03.Username = "user03"
	fmt.Printf("user03:%p\n", user03)
	fmt.Printf("user03:%#v\n", user03)
}

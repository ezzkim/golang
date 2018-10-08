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
	/*
		var user User
		user.Username = "laozhang"
		user.Age = 12
		user.Sex = "male"
		user.AvatarUrl = "www.baidu.com"
	*/
	//var user User = User {  //ok
	user := User{
		Username:  "laozhang",
		Age:       12,
		Sex:       "male",
		AvatarUrl: "www.baidu.com",
	}

	var user2 User

	fmt.Printf("%v\n", user)
	fmt.Printf("user2 = %+v\n", user2)

	var user3 *User = &User{
		Username:  "laozhang",
		Age:       12,
		Sex:       "male",
		AvatarUrl: "www.baidu.com",
	}
	fmt.Printf("user3 = %+v\n", *user3)
}

package main

import (
	"fmt"
	"github.com/study/golang/day12/user"
)

func main() {
	/*
	var user user.User = user.User{
		Username:  "laozhang",
		Age:       12,
		Sex:       "male",
		AvatarUrl: "www.baidu.com",
	}
	*/
	//fmt.Printf("user = %#v\n", user)
	u := user.NewUser("zhang", "female", 25, "www.google.com")

	fmt.Printf("u = %#v\n", u)
}
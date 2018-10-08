package main

import (
	"encoding/json"
	"fmt"
)

//大写可外部访问， 小写不可以
type User struct {
	Username string `json:"xxx",db:"user_name"`
	Sex      string `json:"sex"`
	Score    float32
	age      int32 // no join serial
	number   int32 `josn:"noAccess"`
}

func main() {
	user := &User{
		Username: "user01",
		Sex:      "man",
		Score:    99.5,
		age:      18,
	}

	//json only access the struct can access attribute
	data, _ := json.Marshal(user) // json serial user
	fmt.Printf("json str : %s\n", string(data))
	fmt.Println("----------")
	test()
}

//var str = ` {"xxx":"user06","sex":"female","Score":98.5}`
var str = ` {"xxx":"user06","sex":"female","Score":98.5, "noAccess":123}`

func test() {
	var user User
	json.Unmarshal([]byte(str), &user)
	fmt.Printf("user=%#v\n", user)
}

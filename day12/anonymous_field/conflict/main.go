package main

import (
	"fmt"
)

type Email struct {
	account string
	CreateTime string
}

//inner struct
type Address struct {
	Provice string
	City    string
	CreateTime string
}

type User struct {
	Username string
	Sex      string
	int
	string
	address Address
}

type User1 struct {
	City     string
	Username string
	Sex      string
	*Address // aynmous struct
}

type User2 struct {
	Username string
	Sex      string
	*Address // aynmous struct
}

type User01 struct {
	City string
	Username string
	Sex string
	*Address
	*Email
}

func main() {
	var user User
	user.Username = "ershazi"
	user.Sex = "man"
	user.int = 10
	user.string = "hello"
	user.address.Provice = "shanghai"
	user.address.City = "qingpu"
	fmt.Printf("user = %#v\n", user)

	up := &User{
		Username: "user01",
		Sex:      "man",
		int:      100,
		string:   "hello",
		address: Address{
			Provice: "henan",
			City:    "zhengzhou",
		},
	}

	fmt.Printf("up = %#v\n", up)

	////////////aynmous fuzhi method 1///////
	var u2 User2
	u2.Username = "u2"
	u2.Sex = "female"

	//method 1 ok
	/*
		u2.Address = &Address {
			Provice : "bj",
			City : "haiding",
		}*/

	//method 2
	u2.Address = new(Address) // must fenbei firstly
	u2.Provice = "bj2"
	u2.City = "xianghe"

	fmt.Printf("u2=%#v, addr=%#v\n", u2, u2.Address)
	//fmt.Printf("u2.addres = %v\n", u2.Address)

	fmt.Println("------------")
	testCity()

	fmt.Println("-------------")
	test02()
}

func testCity() {
	var user User1
	user.Username = "laozhang"
	user.Sex = "man"
	user.City = "bj"
	fmt.Printf("user=%#v", user)

	user.Address = new(Address) // must do
	user.Address.City = "bj-inner"
	fmt.Printf("user=%#v city=%s\n", user, user.Address.City)
}

func test02() {
	var user User01
	user.Username = "user01"
	user.Sex = "man"
	user.City = "shanghai"
	user.Address = new(Address)
	user.Email = new(Email)

	fmt.Printf("user=%#v\n", user)
	user.City = "shanghai-01"
	fmt.Printf("user=%#v city of address=%s\n", user, user.Address.City)

	//user.CreateTime = "001" // error
	user.Address.CreateTime = "001"
	user.Email.CreateTime = "002"
	fmt.Printf("user=%#v createtime of address=%s\n", user, user.Address.CreateTime)
	fmt.Printf("user=%#v createtime of email=%s\n", user, user.Email.CreateTime)
}

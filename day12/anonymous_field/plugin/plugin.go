package main

import "fmt"

type Address struct {
	City     string
	Province string
}

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
	address   *Address
}

type User2 struct {
	Username  string
	Sex       string
	*Address
}

func main() {
	var addr Address = Address{
		City:     "shanghai",
		Province: "qingpu",
	}
	fmt.Printf("addr = %#v\n", addr)

	var user *User = new(User)
	user.Username = "user01"
	user.Sex = "man"
	user.Age = 19
	user.AvatarUrl = "jellp"
	//user.address = &addr // ok
	user.address = &Address {
		City:     "shanghai",
		Province: "qingpu",
	}

	fmt.Printf("user = %#v\n", user)

	fmt.Println("--------------")
	anonymousStruct()
}

func anonymousStruct() {
	var user *User2 = new(User2)
	user.Username = "user02"
	user.Sex = "man"
	
	//the first method
	user.Address = &Address {
		City:     "shanghai",
		Province: "qingpu",
	}

	//the second method
	//just modify, must call above code
	user.Province = "xidan"
	user.City = "beijing"

	fmt.Printf("user=%#v, addr=%#v\n", user, user.Address)
}
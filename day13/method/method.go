package main

import (
	"fmt"
)

type People struct {
	Name    string
	Country string
}

func (p *People) Print() {
	fmt.Printf("name[%s] country[%s]\n", p.Name, p.Country)
}

//below different is very important
//func (p *People) Set(name string, country string) {
func (p People) Set(name string, country string) { //just modify the copy object
	p.Name = name
	p.Country = country
}

type Integer int

func (i Integer) Print() {
	fmt.Println(i)
}

func main() {
	var p1 People = People{
		Name:    "laozhang",
		Country: "china",
	}
	p1.Print() // auto convert to pointer
	p1.Set("laosun", "taiwan") 
	//p1.Print() 
	(&p1).Print()

	var a Integer
	a = 100
	a.Print()

	var b int = 200
	a = Integer(b)
	a.Print()
}

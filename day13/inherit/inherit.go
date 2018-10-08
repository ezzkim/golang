package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Sex  string
}

func (a *Animal) Talk() {
	fmt.Printf("I talk, I am %s\n", a.Name)
}

type Dog struct {
	Feet string
	//sAnimal  // anonymous object, value type. copy
	*Animal
	*PuruAnimal
}

func (d *Dog) Eat() {
	fmt.Println("dog is eat")
}

func (d *Dog) Talk() {
	fmt.Printf("Self talk, I am %s\n", d.Name)
}

type PuruAnimal struct {

}

func (p *PuruAnimal) Talk() {
	fmt.Println("puru animal talk")
}

func main() {
	var d *Dog = &Dog{
		Feet: "four feet",
		Animal: &Animal{
			Name: "habaogou",
			Sex: "male",
		},
	}
	d.Eat() // self method
	d.Talk() // inherit from Animal
	d.Animal.Talk()
	d.PuruAnimal.Talk()
}

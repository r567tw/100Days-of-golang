package main

import "fmt"

type NewPerson struct {
	Name string
	Age  int
}

func (p *NewPerson) setAge(newAge int) {
	p.Age = newAge
}

func (p NewPerson) setAgeFail(newAge int) NewPerson {
	p.Age = newAge
	return p
}

func main() {
	var i int
	i = 12

	var jimmy NewPerson
	jimmy = NewPerson{
		Name: "Jimmy",
		Age:  29,
	}

	fmt.Println(&i)

	jimmy.setAge(30)
	fmt.Println(jimmy.Age)

	jimmy.setAgeFail(40)
	fmt.Println(jimmy.Age)
}

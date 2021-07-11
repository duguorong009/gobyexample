package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(Person{name: "prova", age: 15})
	fmt.Println(Person{"JIN", 20})
	fmt.Println(Person{name: "Jack"})
	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}

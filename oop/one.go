package oop

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	id    int
	score int
}

func (p Person) showInfo() {
	fmt.Printf("My name is %s, age is %d\n", p.name, p.age)
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (s *Student) showInfo() {
	fmt.Printf("Name is %s\n", s.Person.name)
	fmt.Printf("Age is %d\n", s.Person.age)
	fmt.Printf("Id is %d\n", s.id)
	fmt.Printf("Score is %d\n", s.score)
}

func Test() {
	person := Person{"mike", 18}
	person.showInfo()
	person.SetAge(100)
	person.showInfo()

	s := Student{
		person,
		123456,
		100,
	}
	s.showInfo()
}

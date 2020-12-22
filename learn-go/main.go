package main

import "fmt"

type Human interface {
	say()
}

type Male interface {
	Human
	MalePlay()
}

type Female interface {
	Human
	FemalePlay()
}

type Student struct {
	Name string
}

func (s Student) say() {
	fmt.Printf("%s是个human", s.Name)
}
func (s Student) MalePlay() {
	fmt.Printf("%s是个male", s.Name)
}
func NewStudent() Human {
	return Student{Name: "Devil"}
}
func main() {
	student := NewStudent()
	male, ok := student.(Male)
	if ok {
		male.MalePlay()
		male.say()
	}
}

package main

import "fmt"

type Human struct {
	Age  int
	Name string
}
type A struct {
	Age int
}

func (a A) SayHello() {
	fmt.Println("Hello from A")
}

func (h Human) SayHi() {
	fmt.Println("Hi from Human")
}

func (h Human) SayBye() {
	fmt.Println("Bye from Human")
}

func NewHuman(age int, name string) Human {
	return Human{
		Age:  age,
		Name: name,
	}
}

type Action struct {
	Human
	A
}

func (a Action) SayHi() {
	fmt.Println("Hi from Action")
}

func NewAction(h Human) Action {
	return Action{
		Human: h,
	}
}

func main() {
	human := NewHuman(25, "Bob")
	action := NewAction(human)
	action.A = A{Age: 20}

	action.SayHi()
	action.SayBye()
	action.SayHello()

	fmt.Println(human, action)
	fmt.Println(action.A.Age, action.Name)

}

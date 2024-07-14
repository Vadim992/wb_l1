package main

import "fmt"

type Adapter interface {
	Operation()
}

type Adaptee struct {
}

func NewAdaptee() *Adaptee {
	return &Adaptee{}
}

func (a *Adaptee) AdapteeOperation() {
	fmt.Println("Adaptee opration")
}

type ConcreteAdapter struct {
	*Adaptee
}

func (c *ConcreteAdapter) Operation() {
	c.AdapteeOperation()
}
func NewConcreteAdapter(a *Adaptee) *ConcreteAdapter {
	return &ConcreteAdapter{a}
}

func NewAdapter(a *Adaptee) Adapter {
	return NewConcreteAdapter(a)
}

func main() {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	adapter.Operation()
}

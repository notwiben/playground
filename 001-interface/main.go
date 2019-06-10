package main

import "fmt"

type I interface {
	M()
	N() string
}

type T struct {
	name     string
	lastName string
}

func (t T) M() {
	t.name = "asd"
}

func (t T) N() string {
	return t.name
}

func Hello(i I) {
	i.M()
	fmt.Printf("Hi, my name is %s\n", i.N())
}

func main() {
	Hello(T{})
}

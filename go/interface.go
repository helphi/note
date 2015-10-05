package main

import "fmt"

func main() {
	var a Integer = 1
	var b LessAdder = &a
	b.Add(1)
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b.(*Integer))
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

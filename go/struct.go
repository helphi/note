package main

import ."fmt"
import "strconv"

func main() {
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{*new(Shape), 0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200, Shape: Shape{10, "rect"}}

	Printf("%v: %v\n", rect1.GetId(), rect1.Area())
	Printf("%s: %v\n", rect2.GetName(), rect2.Area())
	Printf("%s: %v\n", rect3.GetName(), rect3.Area())
	Printf("%s: %v\n", rect4.GetName(), rect4.Area())

	tri1 := Triangle{new(Shape), 4, 5, 6}
	tri2 := Triangle{&Shape{20, "tri2"}, 7, 8, 9}

	Printf("%s: %v\n", tri1.GetName(), tri1.GetCircumference())
	Printf("%s: %v\n", tri2.GetName(), tri2.GetCircumference())
}

type Shape struct {
	id int
	name string
}

func (s Shape) GetId() int {
	return s.id
}

func (s Shape) GetName() string {
	return s.name
}

type Rect struct {
	Shape
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return (*r).width * r.height
}

func (r Rect) GetName() string {
	return r.Shape.GetName() + strconv.Itoa(r.Shape.id)
}

type Triangle struct {
	*Shape
	x, y, z float64
}

func (t Triangle) GetCircumference() float64 {
	Println("in GetCircumference:" + t.GetName())
	return t.x + t.y + t.z
}


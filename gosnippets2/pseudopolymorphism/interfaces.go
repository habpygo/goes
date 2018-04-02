package main

import "fmt"

type Rectangle struct {
	width  int
	length int
}

type Circle struct {
	Radius float32
	Pi     float32
}

type Sphere struct {
	Circle
}

func (r Rectangle) Surface() int {
	return r.length * r.width
}

func (c Circle) Surface() float32 {
	return c.Pi * c.Radius * c.Radius
}

func (s Sphere) Surface() float32 {
	return 4 * s.Circle.Pi * s.Circle.Radius * s.Circle.Radius
}

func main() {

	newRect := Rectangle{width: 25, length: 50}
	fmt.Println("Surface is: ", newRect.Surface())

	newCircle := Circle{Radius: 34, Pi: 3.14159265359}
	fmt.Println("Surface of circle is: ", newCircle.Surface())

	newSphere := Sphere{
		Circle: Circle{
			Pi:     3.14159265359,
			Radius: 34,
		},
	}
	fmt.Println("Surface of sphere is: ", newSphere.Surface())

}

package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func (s *Square) Perimeter() float64 {
	return 4 * s.Side
}

func (s *Square) String() string {
	return fmt.Sprintf("Square. side=%f", s.Side)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle. width=%f, height=%f", r.Width, r.Height)
}

func main() {
	square := &Square{Side: 5}
	rectangle := &Rectangle{Width: 4, Height: 8}
	shapes := []Shape{square, rectangle}
	for _, shape := range shapes {
		fmt.Printf("%s => area=%f, perimeter=%f\n",
			shape, shape.Area(), shape.Perimeter())
	}

	// circle := &Circle{radius: 5}
	// triangle := &Triangle{base: 3, height: 7}
}

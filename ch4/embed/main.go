package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	w1 := Wheel{
		Circle: Circle{
			Point:  Point{X: 4, Y: 2},
			Radius: 6,
		},
		Spokes: 4, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", w1)

	w.X = 42

	w1.Circle.Point.X = 20

	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", w1)
}

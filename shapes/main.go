package main

import "fmt"

type rectangle struct {
	nm string
	l  float64
	b  float64
}

type square struct {
	nm   string
	side float64
}

type shape interface {
	getArea() float64
}

func main() {

	rt := rectangle{nm: "rectangle", l: 10, b: 5}
	sq := square{nm: "square", side: 10}

	printArea(rt.nm, rt)
	printArea(sq.nm, sq)

}

func (r rectangle) getArea() float64 {
	return 2 * (r.l + r.b)
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(nm string, s shape) {
	fmt.Println("Area of the", nm, ": ", s.getArea())
}

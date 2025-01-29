package main

type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	println(s.area())
}
func main() {
	s := square{
		length: 5,
		width:  5,
	}
	c := circle{
		radius: 5,
	}
	info(s)
	info(c)
}

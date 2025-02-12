package main

import "fmt"

func main() {
	x := 42
	y := &x
	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
}

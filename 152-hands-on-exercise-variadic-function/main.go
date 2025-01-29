package main

import "fmt"

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
func main() {
	defer fmt.Println("This is a defer statement")
	x := foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	y := bar([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	println(x, y)
}

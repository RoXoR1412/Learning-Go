package main

import "fmt"

func main() {
	//Array
	ar := [5]int{4, 6, 7, 8, 9}
	for _, v := range ar {
		fmt.Printf("Value %v\t Type %T\n", v, v)
	}
}

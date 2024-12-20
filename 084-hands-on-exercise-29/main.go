package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for y := 5; y > i; y-- {
			fmt.Print(y)
		}
		fmt.Println()
	}
}

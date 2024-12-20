package main

import "fmt"

func main() {
	i := 0
	fmt.Printf("odd numbers are ")
	for i <= 10 {
		if i%2 == 0 {
			i++
			continue
		} else if i%2 != 0 {
			fmt.Printf("%v\t", i)
		}
		i++
	}
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 42; i++ {
		x := rand.Intn(5)
		switch {
		case x == 0:
			{
				fmt.Printf("Iteration = %v and Value of X = %v \n", i, x)
			}
		case x == 1:
			{
				fmt.Printf("Iteration = %v and Value of X = %v \n", i, x)
			}
		case x == 2:
			{
				fmt.Printf("Iteration = %v and Value of X = %v \n", i, x)
			}
		case x == 3:
			{
				fmt.Printf("Iteration = %v and Value of X = %v \n", i, x)
			}
		case x == 4:
			{
				fmt.Printf("Iteration = %v and Value of X = %v \n", i, x)
			}
		}
	}
}

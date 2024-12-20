package main

import (
	"fmt"

	"math/rand"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("Value of x is %v\t", x)
	fmt.Printf("Value of y is %v\t\n", y)
	if x < 4 && y < 4 {
		fmt.Println("X & Y both are less than four")
	} else if x > 6 && y > 6 {
		fmt.Println("X & Y are both greater than six")
	} else if (x >= 4 && y >= 4) && (x <= 6 && y <= 6) {
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("None of the cases were met")
	}
}

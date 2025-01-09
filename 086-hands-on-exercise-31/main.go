package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	//for i, v := range m {
	//	fmt.Printf("Key= %v \t Value=%v \n", i, v)
	//}
	age, ok := m["Q"]
	if ok {
		fmt.Printf("age %v", age)
	} else {
		fmt.Println("Q is not in the map")
	}
}

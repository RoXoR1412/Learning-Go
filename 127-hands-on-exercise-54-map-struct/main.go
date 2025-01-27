package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName:       "James",
		lastName:        "Bond",
		iceCreamFlavors: []string{"Chocolate", "Vanilla"},
	}
	p2 := person{
		firstName:       "Miss",
		lastName:        "Moneypenny",
		iceCreamFlavors: []string{"Strawberry", "Mint"},
	}
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	//fmt.Println(m)
	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.iceCreamFlavors {
			fmt.Println(v.firstName, v.lastName, v2)
		}
	}
}

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
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.iceCreamFlavors)
	fmt.Println(p2.iceCreamFlavors)

	for _, v := range p1.iceCreamFlavors {
		fmt.Println(p1.firstName, "'s favourite ice-cream is ", v)
	}
	for _, v := range p2.iceCreamFlavors {
		fmt.Println(p2.firstName, "'s favourite ice-cream is ", v)
	}

}

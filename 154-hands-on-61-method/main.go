package main

type person struct {
	first string
	age   int
}

func (p person) speak() {
	println("I am", p.first, "and I am", p.age, "years old.")
}
func main() {
	p1 := person{
		first: "James",
		age:   32,
	}
	p1.speak()
}

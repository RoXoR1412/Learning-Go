package main

import "fmt"

func main() {
	ma := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}
	ma[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	delete(ma, "no_dr")
	for i, v := range ma {
		fmt.Printf("%v\t%v\n", i, v)
	}
}

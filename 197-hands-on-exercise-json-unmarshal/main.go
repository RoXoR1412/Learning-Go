package main

import (
	"encoding/json"
	"fmt"
)

type Person []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	bs := []byte(s)
	var people Person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("Person", i+1)
		fmt.Println("\tFirst:", people[i].First)
		fmt.Println("\tLast:", people[i].Last)
		fmt.Println("\tAge:", people[i].Age)
		for j := 0; j < len(people[i].Sayings); j++ {
			fmt.Println("\t\t", people[i].Sayings[j])
		}
	}
	//fmt.Println(people)

}

package main

import "fmt"

type Name struct {
FirstName string
LastName string
}
func main() {
	name1 := Name {
		FirstName: "Steve",
		LastName: "Jobs",
	}
	name2 := Name {
		FirstName: "Steve",
		LastName: "Jobs",
	}
	if (name1 == name2) {
		fmt.Println("name1 and name2 are equal")
	} else {
	fmt.Println("name1 and name2 are not equal")
	}

	name3:= Name {
		FirstName: "Steve",
		LastName: "Jobs",
	}
	name4:= Name {
		FirstName: "Lisa",
		LastName: "Roy",
	}
	if (name3 == name4) {
		fmt.Println("name3 and name4 are equal")
	} else {
	fmt.Println("name3 and name4 are not equal")
	}
}

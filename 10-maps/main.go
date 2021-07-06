package main

import "fmt"

func main() {
	studentsAge := map[string]int{
		"jonh":  42,
		"peter": 19,
	}

	fmt.Println(studentsAge)

	map2 := make(map[string]int)
	map2["test1"] = 1
	map2["test2"] = 2

	fmt.Println(map2)

	fmt.Println("Peter's age is:", studentsAge["peter"])

	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

	delete(map2, "test2")
	fmt.Println(map2)
}

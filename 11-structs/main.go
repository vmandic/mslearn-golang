package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	ID      int    `json:"id"`
	Address string `json:"addr"`
	Info    Info   `json:"info"`
}

type Info struct {
	PhoneNumber int `json:"phone"`
}

func main() {
	var p1 Person
	p1.Name = "John"
	p1.ID = 1
	p1.Address = "Street 123"
	p1.Info.PhoneNumber = 2226666
	fmt.Println(p1)

	p2 := Person{"Jack", 2, "Somewhere 23123", Info{555123}}
	fmt.Println(p2)

	p3 := Person{Name: "Mary", ID: 3, Address: "Yup3333"}
	p3.Info.PhoneNumber = 555000
	fmt.Println(p3)
	fmt.Println(&p3)

	p4 := &p3
	fmt.Println(p4)

	p4.Name = "Newguy"

	fmt.Println(p4)
	fmt.Println(*p4)
	fmt.Println("p3 after changing p4 created from &p3", p3)

	changeNameCopy(p1)
	fmt.Println("after change name copy", p1)

	changeNamePointer(&p1) // send by address
	fmt.Println("after change name pointer", p1)

	data1, _ := json.Marshal(p1)
	fmt.Printf("JSON: %s\n", data1)

	var p1Deserialized Person
	json.Unmarshal(data1, &p1Deserialized)
	fmt.Println("JSON deserialized:", p1Deserialized)
}

func changeNameCopy(p Person) {
	p.Name = "changed!!!"
}

func changeNamePointer(p *Person) { // NOTE dereference value from address
	p.Name = "changed!!!"
}

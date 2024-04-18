package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Position string `json:"position"`
}

func main() {

	data, err := os.ReadFile("employees.json")
	if err != nil {
		log.Fatal(err)
	}

	var persons []Person
	err = json.Unmarshal(data, &persons)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(persons); i++ {
		if persons[i].Id == 2 {
			persons[i].Age = 50
		}
	}

	persons = append(persons, Person{
		Id:       6,
		Name:     "Bob Robin",
		Age:      30,
		Position: "Software Engineer",
	})

	body, err := json.Marshal(persons)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("employees_new.json")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(file.Name(), body, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(persons)

}

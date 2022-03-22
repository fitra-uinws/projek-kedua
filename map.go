package main

import "fmt"

func main() {
	var person map[string]string

	person = map[string]string{}
	person["name"] = "Hacktiv8"
	person["age"] = "8"
	person["address"] = "Jalan loren ipsum"

	fmt.Println("nama:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	person = map[string]string{
		"name":    "Hacktiv8",
		"age":     "9",
		"address": "Jalan Loren ipsum 1",
	}
	fmt.Println("nama:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	fmt.Println("==================LOOPING==============")
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

	fmt.Println("==================delete map==============")
	delete(person, "name")
	fmt.Println(person)

	fmt.Println("==================detect map==============")
	value, exists := person["name"]

	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key already deleted")
	}

	fmt.Println("==================slice and map==============")

	people := []map[string]string{
		{"name": "Hactiv8", "age": "6"},
		{"name": "Hactivkids", "age": "2"},
		{"name": "Kode", "age": "5"},
	}

	for key, person := range people {
		fmt.Printf("index:%d, name: %s, age: %s \n", key, person["name"], person["age"])
	}

}

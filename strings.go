package main

import "fmt"

func main() {
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index : %d, byte: %d \n", i, city[i])
	}

	cityByte := []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(cityByte))

	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte lenght ==> %d ", len(str1))
	fmt.Printf("str1 byte lenght ==> %d ", len(str2))
}

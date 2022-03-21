package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var stringss = [3]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%#v \n", numbers)
	fmt.Printf("%#v \n", stringss)

	var fruits = [3]string{"Apel", "Pisang", "Mangga"}

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Mango"

	fmt.Printf("%#v \n", fruits)

	fmt.Println("====Looping Array 1====")
	for i, v := range fruits {
		fmt.Printf("Index: %v, Value: %s \n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	fmt.Println("====Looping Array 2====")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %v, Value: %s \n", i, fruits[i])
	}

	fmt.Println(strings.Repeat("#", 25))

	fmt.Println("====Array Multi dimensi====")
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}
	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	fmt.Println("====slice====")
	sliceFruits := []string{"Apple", "Banana", "Mango"}

	_ = sliceFruits

	sliceFruitsMake := make([]string, 3)
	sliceFruitsMake[0] = "Apple"
	sliceFruitsMake[1] = "Banana"
	sliceFruitsMake[2] = "Mango"

	fruitsBaru := []string{"Jambu", "Anggur", "Jeruk"}

	//sliceFruitsMake = append(sliceFruitsMake, "Apple2", "Banana2", "Mango2")
	sliceFruitsMake = append(sliceFruitsMake, fruitsBaru...)

	fmt.Printf("%#v \n", sliceFruitsMake)

	fmt.Println("====copy slice====")
	nn := copy(sliceFruitsMake, fruitsBaru)
	fmt.Printf("sliceFruitsMake %#v \n", sliceFruitsMake)
	fmt.Printf("fruitsBaru %#v \n", fruitsBaru)
	fmt.Printf("copied elements %#v \n", nn)

	fmt.Println("====slicing====")

	fruits1 := []string{"Apple", "Banana", "Mango", "Jambu", "Anggur", "Jeruk"}
	fruits2 := fruits1[1:4]
	fmt.Printf("%#v \n", fruits2)
	fruits3 := fruits1[2:]
	fmt.Printf("%#v \n", fruits3)
	fruits4 := fruits1[:3]
	fmt.Printf("%#v \n", fruits4)

	fmt.Println("====slice Append with slice====")
	fruits5 := append(fruits1[:3], "orange")
	fmt.Printf("%#v \n", fruits5)

	//backing slice
	//header slice
	//1. alamat memory
	//2. panjang slice bisa gunakan len
	//3. kapasitas dari slice  bisa gunakan cap

	fruits6 := fruits1[2:4]
	fruits6[0] = "Rambutan"

	fmt.Println("fruits1 ==>", fruits1)
	fmt.Println("fruits6 ==>", fruits6)

	fmt.Println("====Capacity (cap)====")
	fmt.Println("Fruits1 cap:", cap(fruits1))
	fmt.Println("Fruits1 len:", len(fruits1))
	fmt.Println(strings.Repeat("#", 25))

	fmt.Println("fruits6 cap:", cap(fruits6))
	fmt.Println("fruits6 len:", len(fruits6))
	fmt.Println(strings.Repeat("#", 25))

	fmt.Println("====new Backing Array====")
	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "suzuki"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)
}

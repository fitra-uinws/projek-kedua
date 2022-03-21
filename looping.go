package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	i := 0
	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}

	i = 0
	for {
		fmt.Println("Angka", i)
		i++

		if i == 3 {
			break
		}
	}

	//i = 0
	fmt.Println("if continue break")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if 1 > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	fmt.Println("=============latihan==================")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}
	fmt.Println("=============outerLoop==================")
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("perulangan ke -", i+1)

		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Println(j, "")
		}
		fmt.Println()
	}
}

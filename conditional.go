package main

import "fmt"

func main() {
	currentYear := 2022

	if age := currentYear - 1993; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	score := 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

	score = 6
	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 7 && score > 3):
		fmt.Println("Not bad")
	default:
		fmt.Println("Study harder")
		fmt.Println("you need to learn more")
	}

	//falltrought case
	score = 4
	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8 && score > 3):
		fmt.Println("Not bad")
		fallthrough
	case score < 5:
		fmt.Println("it s ok, but please study harder")
	default:
		fmt.Println("Study harder")
		fmt.Println("you need to learn more")
	}

	//nested condition
	score = 10
	if score > 7 {
		switch score {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice")
		}
	} else {
		if score == 5 {
			fmt.Println("Not Bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder")
			}
		}

	}
}

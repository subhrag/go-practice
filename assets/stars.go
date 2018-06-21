package main

import "fmt"

func displayStars(row int) {
	for i := 0; i < row; i++ {
		for j := 0; j < i+i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}

}

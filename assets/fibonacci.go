package main

import "fmt"

func calculateFib() {

	var number = 10
	var result int
	var firstNumber, secondNumber int = 0, 1
	fmt.Println("Here the febonacci series is ")
	for i := 0; i < number; i++ {
		result = firstNumber + secondNumber
		firstNumber = secondNumber
		secondNumber = result
		fmt.Print(result, " - ")
	}
}

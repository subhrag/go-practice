package main

import "fmt"

func main() {
	fmt.Println("hi")
	calculateFib()
	displayStars(5)
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Area of shape", getShapeArea(rect))
}

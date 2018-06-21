package main

type Shape interface {
	shapeArea() float64
}

type Rectangle struct {
	width, height float64
}

func (rectangle Rectangle) shapeArea() float64 {
	return rectangle.width * rectangle.height
}

func getShapeArea(shape Shape) float64 {
	return shape.shapeArea()
}

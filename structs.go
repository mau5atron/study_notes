package main

import "fmt"

// Structs

func main(){
	// Define a rectangle
	rect1 := Rectangle {
		leftX: 0, 
		topY: 50, 
		height: 10, 
		width: 10,
	}

	// leave off attribute names if we know the order 
	// rect1 := Rectangle{0, 50, 10, 10}

	// we access the values with the dot operator
	fmt.Println("Rectangle is", rect1.width, "wide")
	// call the area method for Rectangle
	fmt.Println("Area of the rectangle =", rect1.area())
}

// we can define our own types ising struct 
// this defines what the heck a rectanlge is with attributes
type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}

// We can define methods for our rectangle by adding the receiver 
// rect *Rectangle between func and the function name so we can
// call it with the dot operator
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height  
}
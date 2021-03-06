package main

import "fmt"

func main(){
	// A constant is a variable with a value that cannot be changed
	const pi float64 = 3.14159265359

	// You can declare multiple variables like this
	var (
		varA = 2
		varB = 3
	)

	fmt.Println(varA, varB) // => 2 3

	// Strings are a series of characters between "" or ''
	var myName string = "Gabriel Betancourt"

	// Get string length
	fmt.Println(len(myName)) // => 18

	// You can concatenate strings with + 
	fmt.Println(myName + " is a robot") // => Gabriel Betancourt is a robot

	// Strings between "" can contain escape symbols
	fmt.Println("I like \n \n") // two new lines
	fmt.Println("Newlines") // => 
	/* => 
	I like 


	Newlines
	*/

	// Booleans can be either true or false
	var isOver40 bool = true

	fmt.Println(isOver40) // => true

	// Printf is used for format printing (%f is for floats)
	fmt.Printf("%f \n", pi) // => 3.141593 

	// You can also define the decimal precision of a float
	fmt.Printf("%.3f \n", pi) // => 3.142

	// %T prints the data type
	fmt.Printf("%T \n", pi) // => float64

	// %t prints booleans
	fmt.Printf("%t \n", isOver40) // => true

	// %d is used for integers
	fmt.Printf("%d \n", 100) // => 100

	// %b prints in binary
	fmt.Printf("%b \n", 100) // => 1100100 

	// %c prints the character associated with a keycode
	fmt.Printf("%c \n", 44) // => ,

	// %x prints in hexcode
	fmt.Printf("%x \n", 17) // => 11

	// %e prints in scientific notation
	fmt.Printf("%e \n", pi) // => 3.141593e+00

}


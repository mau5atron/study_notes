// Go programs start with a package declaration to make the code resuable

package main 


// Import allows the use of other packages 
// The format package provides formatting for input and output
import "fmt"

// A comment
/*
	Multiline comment
*/ 

// Functions start with func and surround the code with {}
// main is the function that is executed when you execute your program

func main(){

	// Println is a function in the fmt package that outputs a string, 
	// which is surrounded by double quotes and a newline to the screen

	fmt.Println("Hello World")

	// You can obtain a description of a function by typing 
	// ****** godo fmt Println

	// Go programs are executed by typing into terminal 
	// go run herewego.go

	// Variables are statically typed, which means their type can't change
	// varaible names must start withq letter and may contain letters, numbers
	// or the _

	// An int is a positive or negative number without decimals 
	// Version
	// unint8 : unsigned 8-bit integers (0 to 255)
	// unint16 : unsigned 16-bit
	// unint32 : unsigned 32-bit integers (0 to 4294967295)
	// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
	// int8 : signed 8-bit integer (-128 to 127)
	// int16 : signed 16-bit integer (-32768 to 32767)
	// int32 : signed 32-bit integer (-2147483648 to 2147483647)
	// int64 : signed 64-bit integer (-9223372036854775808 to 9223372036854775807)

	var age int64 = 40

	// a float is a number with decimals
	// Versions : float32, float64

	var favNum float64 = 1.61803398875

	// data types do not need to be defined, nor do you have to use semicolons
	// however it is good practice

	randNum := 1;
	fmt.Prinln(randNum)

	// You cannot assign a non-compatible type
	// randNum = "Hello"

	fmt.Println(age, " ", favNum)

	var numOne = 1.000
	var num99 = .999

	// You can perform arthmetic in Println (Note that floats are not accurate)

	fmt.Println(numOne - num99)

	// You can perform arithmetic in Println

	// Artithmetic Operators : +, -, *, /, %

	fmt.Println("6 + 4 = ", 6 + 4)
	fmt.Println("6 - 4 = ", 6 - 4)
	fmt.Println("6 * 4 = ", 6 * 4)
	fmt.Println("6 / 4 = ", 6 / 4)
	fmt.Println("6 % 4 = ", 6 % 4)
}
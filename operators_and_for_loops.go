package main

import "fmt"

func main(){
	// Logical Operators SON !
	fmt.Println("true && false", true && false) // first falsy
	fmt.Println("true || false", true || false) // first truthy
	fmt.Println("!true =", !true) // not true! obviously equals false

	// For Loops
    i := 1

	for i <= 10 {
		fmt.Println(i)
		// shorthand for
		i++
	}

	// Relational Operators include ==, !=, <, >, <=, >=
	
	// You can also define a for loop like this, but you need semicolons
	// j is initialized as 0, if j is less than 5, then increment j
	for j := 0; j < 5; j++{
		fmt.Println(j)
	}
}
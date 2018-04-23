package main

import "fmt"

func main(){
	// If statement
	yourAge := 22

	if yourAge >= 16 {
		fmt.Println("You can drive bro no prob")
	} else {
		fmt.Println("Uuuhh you can't drive")
	}


	// You can use else if to perform different actions, but once a match is reached
	// the rest of the conditions are not checked 
	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can have fun")
	}


	// switch statements are used when you have limited options

	switch yourAge {
		case 16: fmt.Println("Go drive")
		case 18: fmt.Println("You can vote")
		default: fmt.Println("Go have fun")
	}
}
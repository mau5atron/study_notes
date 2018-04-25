package main

import "fmt"

func main(){
	// A map is a collection of key value pairs 
	// Created with varName := make(map[keyType] valueType)
	
	presAge := make(map[string] int)
	presAge["TheodoreRoosevelt"] = 42
	fmt.Println(presAge["TheodoreRoosevelt"]) // => 42

	// Get the number of items in the map
	fmt.Println(len(presAge)) // => 1

	// The size changes when a new item is added
	presAge["John F. Kennedy"] = 43
	fmt.Println(len(presAge)) // => 2

	// We can de1lete by passing the key to delete
	delete(presAge, "John F. Kennedy")
	fmt.Println(len(presAge)) // => 1
	fmt.Println(presAge) // => map[TheodoreRoosevelt:42]
}
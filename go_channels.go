package main 

import (
	"fmt"
	"time"
	"strconv"
)

// Channels
// Channels allow us to pass data between go routines

// initializing some variables
var pizzaNumber = 0
var pizzaName = ""

func makeDough(stringChan chan string){
	// increments pizzaNumber
	pizzaNumber++

	// Converts int into a string
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNumber)
	fmt.Println("Make dough and send for sauce")
	
	// Send the pizzaName onto the channel for the next
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string){
	// Receive the value passed on the channel
	pizza := <- stringChan
	fmt.Println("Add sauce and Send", pizza, "for Toppings")

	// Send the pizzaName onto the channel for the next 
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string){
	// Receive the value passed on the channel
	pizza := <- stringChan
	fmt.Println("add Toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 10)
}

func main(){
	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond * 5000)
	}
}
package main

import(
	"fmt"
	"time"
)

// GO ROUTINES

func count(id int){
	// the for loop reads: I initialize i as 0, then if i is less than 10, I want to increment i
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)

	// This pauses the count function for 1 second to allow other functions to execute
	time.Sleep(time.Millisecond * 1000)
	}
}


func main(){
	// A go routine is a function tha runs in parallel with other functions
	// We define one by using 'go' followed by the function name

	for i := 0; i < 10; i++ {
		go count(i)
	}

	// Wait for the timer to make sure the go routine has time to finish
	// otherwisse the program would end before that happens
}
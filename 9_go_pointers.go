package main

import "fmt"

func main(){
	// We pass  the value of a variable
	x := 0
	changeXVal(x)
	fmt.Println("x =", x) // => x = 0

	// If we pass a reference to the varible we can change the value in a function
	changeXValNow(&x)
	fmt.Println("x =", x) // => x = 2

	// Get the address x points to in memory with &
	fmt.Println("Memory Address for x = ", &x) // => Memory Address for x =  0xc420014080

	// We can also generate a pointer with new
	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y = ", *yPtr) // => y =  100
}

func changeXVal(x int){
	// Has no effect on the value of x in main()
	x = 2
}

// * signals that we are being sent a reference to the value
func changeXValNow(x *int){
	// change the value at the memory address referenced by the pointer
	// * gives us access to the value the pointer points at

	*x = 2 // Store 2 in the memory address x refers to
}

func changeYValNow(yPtr *int){
	*yPtr = 100
}


package main

import "fmt"

func main(){

	listOfNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum: ", addThemUp(listOfNums))

	// get two values from a function
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)

	// send an undefined number of values to a function (Variadic Function)
	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// You can create a function in a function. It has access tot the local 
	// variables of the containing function
	// A function like this with no local variables is a closure

	num3 := 3

	doubleNum := func() int {
		num3 *= 2
		return num3
	}

	fmt.Println(doubleNum());
	fmt.Println(doubleNum());


	// Calling a recursive function
	fmt.Println(factorial(3))

	// Defer executes a function after the inclosing function finishes 
	// defer can be used to KEEP functions together in a logical way
	// in addition, it executes one last time as a cleanup operation
	// Like defer closing a file after it is opened to perform some operations

	defer printTwo()
	printOne()

	// Use recover() to catch a division by 0 error	
	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))
	// We can catch our own errors and recover with panic and recover
	demPanic()
}	



// Functions allow us to reuse code and provide structure 
// func funcName(parameterPassed) returnType
// Functions do not have access to any variables aside from those passed into them

func addThemUp(numbers []float64) float64{
	sum := 0.0

	for _, val := range numbers {
		// shorthand for sum = sum + val
		sum += val
	}
	return sum
}


// Go functions can return multiple values

func next2Values(number int)(int, int){
	return number+1, number+2
}

// You can receive an undefined number of values with args ...int
func subtractThem(args ...int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}






// An example of recursion : Function calls itself
// Factorial(3)
// 3 * factorial(2) == 3 * 2
// 2 * factorial(2) == 2 * 1
// factorial(0) == 1

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)
}

// Used to demonstrate defer

func printOne(){ fmt.Println(1) }
func printTwo(){ fmt.Println(2) }

// If an error occurs we cant catch the error with recover and allow 
// the code to continue to execute 

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}

// Demonstrate how to call panic and handle it with recover

func demPanic(){
	defer func(){
		// If I did not print the message nothing woudld show
		fmt.Println(recover())
	}()
	panic("PANIC")
}

























package main 

import "fmt"

func main(){
	// An array holds a fixed number of values of the same type

	var favNums2[5] float64

	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618


	// You can access the value by supplying the index number
	fmt.Println(favNums2[3]) // => 3.141


	// Another way of ititializing an array
	favNums3 := [5]float64 {1, 2, 3, 4, 5}
	// Iterating through an array (use _ if a value is not used)
	for i, value := range favNums3 {
		fmt.Println(value, i) 
		// => 
		/*
			1 0	
			2 1
			3 2
			4 3
			5 4		
		*/

	}

	// Slices are like arrays but you leave out the size
	numSlice := []int {5, 4, 3, 2, 1}

	// You can create by defining the first index value to take through the last

	numSlice2 := numSlice[3:5] // numSlice == [2, 1]

	fmt.Println("numSlice2[0] = ", numSlice2[0]) // => numSlice2[0] =  2
	// if you do not supply the first index it defaults to 0
	// if you do not supply the last index it defaults to max
	fmt.Println("numSlice[:2] = ", numSlice[:2]) // => numSlice[:2] =  [5 4]
	fmt.Println("numSlice[2:]", numSlice2[2:]) // => numSlice[2:] []
	fmt.Println("numSlice[1:]", numSlice2[1:]) // => numSlice[2:] [1]


	// You can also create an empty slice and define the data type, 
	// length (receive value of 0), capacity (max size)
	numSlice3 := make([]int, 5, 10) 

	// You can copy a slice to another
	copy(numSlice3, numSlice)

	fmt.Println(numSlice3[0]) // => 5

	// Append values to the end of a slice
	// -1 is the last value in the slice
	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3[6]) // => -1 

}
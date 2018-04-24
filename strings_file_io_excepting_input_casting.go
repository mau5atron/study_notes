package main

import(
	"fmt"
	"strings"
	"sort"
	// "os"
	// "log"
	// "io/ioutil"
	// "strconv"
)

func main(){
	// String functions

	sampString := "Hello world"

	// Returns true if phrase exsists in string
	fmt.Println(strings.Contains(sampString, "lo")) // => true

	// returns the index for the match
	fmt.Println(strings.Index(sampString, "lo")) // => 3

	// Returns the number of matches for the string
	fmt.Println(strings.Count(sampString, "l")) // => 3

	// Replaces the first letter with the second as many times as you define it 
	// This line says to replace l with x three times
	fmt.Println(strings.Replace(sampString, "l", "x", 3)) // => Hexxo worxd

	// Return a list separating with the defined separator 
	csvString := "1, 2, 3, 4, 5, 6"
	fmt.Println(strings.Split(csvString, ",")) // => [1  2  3  4  5  6]

	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters:", listOfLetters) // => Letters: [a b c]

	// This also works with numbers as strings
	listOfNumbers := []string{"3", "1", "2"}
	sort.Strings(listOfNumbers)
	fmt.Println("Numbers:", listOfNumbers) // Numbers: [1 2 3]

	listOfNumbers2 := strings.Join([]string{"3", "2", "1"}, ", ")
	fmt.Println(listOfNumbers2) // => 3, 2, 1

}
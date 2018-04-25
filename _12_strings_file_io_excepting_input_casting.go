package main

import(
	"fmt"
	"strings"
	"sort"
	"os"
	"log"
	"io/ioutil"
	"strconv"
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

	
	// FILE I/O

	// Create a file
	file, err := os.Create("samp.txt")

	// Output any errors
	if err != nil {
		log.Fatal(err)
	}

	// Writing a string to a file
	file.WriteString("This is some text my bro dude")

	// Closing a file
	file.Close()

	// Trying to open the file
	stream, err := ioutil.ReadFile("samp.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert into a string
	readString := string(stream)
	fmt.Println(readString)

	// Accepting input
	fmt.Println("What is your name? ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello", name)

	// Casting
	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	// Convert numbers types
	fmt.Println(float64(randInt)) // => 5
	fmt.Println(int(randFloat)) // => 10

	// Convert a string into an int
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt) // => 100

	// Convert a string into a float
	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat) // => 250.5
}
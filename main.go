package main // declare our package name

//IMPORT packages
import (
	// used to work with json data
	"encoding/json"
	"fmt"                // used to print
	function "functions" // import another .go file and declare it as "function". This imports from your $GOPATH/src/FOLDER_NAME
	"reflect"            // package for checking types easily
	"strconv"            // package for converting to and from strings
)

// DECLARE TYPES / VARS / CONSTS ON A GLOBAL LEVEL
type customString string // Our custom type that is of underlying type string

// Go does NOT support objects or classes like some languages you may be used to instead we create our own type of a Struct
type person struct {
	first string
	last  string
	age   int
}

type car struct {
	Year  int
	Color string
	Make  string
}

// to create a method for our struct, we simply create a function that takes a struct as the param
func (p person) birthYear() string { // here we have a birthYear function that takes a person struct and outputs a string value
	return "Person (struct): " + p.first + " " + p.last + ", and they were born in the year: " + strconv.Itoa(2018-p.age)
}

var ( // declare multiple vars
	s1           = "Hello"
	intSlice     = []int{1, 2, 3} // declare an empty int slice (slices are used instead of arrays in GO. Arrays can be used but best practice is to use slices)
	stringslice  = []string{"This", "comes", "from", "a", "slice"}
	stringSlice2 = []string{"I", "am", "also", "from", "a", "slice"}
	multiSlice   = [][]string{stringslice, stringSlice2} // here we create a multi dimentional slice comprised of two other slices of the same type
	c1           = car{
		Year:  1969,
		Color: "red",
		Make:  "Caddy",
	}
	c2 = car{
		Year:  1990,
		Color: "yellow",
		Make:  "Honda",
	}
	cars = []car{c1, c2}
)

const ( // declare constant values
	custom customString = " World"
)

//FUNCTIONS

// func MAIN is the main function that gets run from your main.go file.
func main() {
	s2 := " from another file!"    // declare a local variable
	s3 := s1 + string(custom) + s2 // concat strings AND convert our custom type to a string

	/*
	*	BELOW WE WILL PRINT OUR RESULTS
	 */
	function.ReturnString(s3) // pass string to function from another file that prints

	fmt.Println() // print empty line

	fmt.Println("Slice #1: ", intSlice) //print our slice

	fmt.Println("Appended Slice: ", append(intSlice, 6, 7, 8, 9, 10)) // append number on top of our slice

	fmt.Println("Deleted slice: ", append(intSlice[:0], intSlice[2:]...)) // use the append method to delete the first 2 numbers and print the rest using ...

	fmt.Println("Sliced slice: ", intSlice[0:2]) //slice our slice using the : operator to only print from index 1 - 3
	fmt.Println()

	fmt.Println("Multidimentional slice: ", multiSlice) // print out multidimentional slice
	fmt.Println()

	//MAPS
	// Here we are creating what is known in GO as a map. A map is a key/value pair that is used for easy look up and storing data.
	// first we define the KEY type - in this case we will use strings. The values in this case will ONLY be ints
	// NOTE: the last value in a mpa must have a trailing comma
	m := map[string]int{
		"Life": 42,
	}

	// Notice the value of "Death" does not print. This is because we wrote some idomatic code to check if there is actually a value matching the key.
	// If you were to try printing the m["Death"] value outside this check, it will return 0.
	if v, ok := m["Death"]; ok {
		fmt.Println(v)
	} else { // if the value does no exist we can create it and assign a value.
		m["Death"] = 999
		fmt.Println("Before deleting map value: ", m["Death"])
		delete(m, "Death") // we can delete from a map and even slices by using the built-in delete function
		fmt.Println("Deleted map value ", m["Death"])
		fmt.Println()
	}

	// Here we will define a map that can take different types by using the interface type
	userMap := map[string]interface{}{
		"username": "Jdoe",
		"age":      69,
		"hasACar":  true,
	}

	fmt.Println("What is the answer to life? ", m) // print our int map
	fmt.Println()

	fmt.Println("Who is our user? ", userMap) // print a value from our map
	fmt.Println()

	for k, v := range userMap {
		fmt.Println("UserMap - ", k, v, "(", reflect.TypeOf(v).String(), ")") // print our userMap key, the value and the type of the value
	}
	fmt.Println()

	//PERSON STRUCTS
	// here we will create a new instance of the person struct
	p1 := person{
		first: "Jane",
		last:  "Doe",
		age:   27,
	}

	defer printBirthYear(p1) // Used the defer keyword to run our function at the end of our func main

	forLoop(intSlice) // call our forLoop function to loop over the slice

	func() {
		fmt.Println()
		fmt.Println("I'm here because of a self executing (anonymous) func!")
		fmt.Println()
		fmt.Println("Below is our factorial function")
		factorial(4, 3)
		fmt.Println()
	}()

	structToJSON(cars)
}

// main.go defined functions
// This function will loop over the value (int) we pass in
func forLoop(value []int) {
	for i, v := range value {
		// For every itteration print the type of the value, the index in the slice and it's value
		fmt.Println("IntSlice Index: ", i, "\n", "IntSlice Value: ", v)
	}
}

func printBirthYear(p person) {
	fmt.Println()
	fmt.Println("Person struct being called by defer: ", p.birthYear())
}

// A recursive function that calls itself!
func factorial(factor int, total int) int {
	if total == 0 {
		return 1
	}
	fmt.Println(factor, " * ", total, " = ")
	factor *= total
	fmt.Println(factor)
	total--
	return factor * factorial(factor, total)
}

func structToJSON(c []car) {
	jsonData, err := json.Marshal(c)
	if err != nil { // if error, print the error
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("Struct converted to JSON: ", string(jsonData))

	fmt.Println(reflect.TypeOf(jsonData))
	// jsonToStruct(jsonData)
}

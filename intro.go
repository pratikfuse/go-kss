package main

import "fmt"

// creating a global variable


// a variable that doesn't have a value will be initialized with default value based on type

var i int // in this case 0
var name string // in this case an empty string

// creating a map
var m = map[string] string {}

// creating an array
var arr = []int{1,2,3,4,5,6}


// init is optional and runs before the main function
func init(){
	fmt.Println("from init function")
}


// main function denotes the entry point of the go application
func main(){

	// print to the console with \n at the end
	fmt.Println("value:", i)
	name = "Fusemachines"
	// print formatted string
	fmt.Printf("We love %s !! \n", name)

	// type int will be inferred from the value with the := operator
	age := 20
	fmt.Printf("Age: %d \n", age)

	m["first"] = "First value"
	m["second"] = "Second value"

	// iterate through a map
	for key, value := range m {
		fmt.Printf("%s:%s \n", key, value)
	}

	// iterate through an array
	for index, value := range arr{
		fmt.Printf("%d:%d \n", index, value)
	}


}


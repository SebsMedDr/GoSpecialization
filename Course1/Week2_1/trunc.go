package main

import "fmt"

func main(){
	var float_input float64 //Set variable for float
	fmt.Printf("Input a floating point number:\n") 
	fmt.Scan(&float_input) //Ask user for float

	trunc_int := int(float_input) //Convert float to int
	fmt.Printf("The truncate integer of the input floating point number is: %v\n",trunc_int) //Print converted float
}
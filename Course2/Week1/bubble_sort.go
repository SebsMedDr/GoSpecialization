package main

import "fmt"
import "strconv"

func Swap(integers []int, j int){
	int1 := integers[j]
	int2 := integers[j+1]
	integers[j] = int2
	integers[j+1] = int1
}

func BubbleSort(integers []int){
	int_len := len(integers)
	for i := 0; i < int_len - 1; i++{
		for j := 0; j < int_len - i - 1; j++{
			if integers[j] > integers[j+1]{
				Swap(integers,j)
			}
		}
	}
}

func main(){

	integers := []int{}
	var integer_str string

	fmt.Print("Please enter 10 integers.\n")
	for i:=0; i<10; i++{
		fmt.Printf("Enter an integer: ")
		fmt.Scan(&integer_str)

		integer, err := strconv.Atoi(integer_str)
		if err == nil{
			integers = append(integers, integer)
		}else{
			fmt.Printf("The format of the input cannot be converted to an integer. Please try something else.\n")
			continue
		}
	}

	BubbleSort(integers)

	fmt.Printf("Printing the sorted slice.\n")
	fmt.Println(integers)
	fmt.Printf("\n\n")
}
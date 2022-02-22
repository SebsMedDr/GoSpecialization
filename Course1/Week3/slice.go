package main

import "fmt"
import "strconv"

func main(){

	integers := make([]int,3)
	integers2 := make([]int,3)
	var integer_str string
	var index int = 0

	fmt.Printf("Please enter integers or X to terminate the program. \n")
	for {
		fmt.Printf("Enter an integer: ")
		fmt.Scan(&integer_str)
		
		if integer_str == "X"{
			fmt.Printf("Terminating the program.\n")
			break
		}

		integer, err := strconv.Atoi(integer_str)
		if err == nil{
			if index < 3{
				integers2[index] = integer
				index += 1
			}else{
				integers = append(integers, integer)
				index += 1
			}
		}else{
			fmt.Printf("The format of the input cannot be converted to an integer. Please try something else.\n")
			continue
		}
		
		if index < 4{
			copy(integers,integers2)
		}

		fmt.Printf("Sorting the slice.\n")
		int_len := len(integers)
		for i := 0; i < int_len - 1; i++{
			for j := 0; j < int_len - i - 1; j++{
				if integers[j] > integers[j+1]{
					int1 := integers[j]
					int2 := integers[j+1]
					integers[j] = int2
					integers[j+1] = int1
				}
			}
		}

		fmt.Printf("Printing the sorted slice.\n")
		fmt.Println(integers)
		fmt.Printf("\n\n")
	}

}
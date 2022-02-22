package main

import "fmt"
import s "strings"

func main(){

	var find_ian_str string

	fmt.Printf("Please enter a string: ")
	fmt.Scan(&find_ian_str)
	find_ian_str = s.ToLower(find_ian_str)

	if s.HasPrefix(find_ian_str,"i") && s.HasSuffix(find_ian_str,"n") && s.Contains(find_ian_str,"a"){
		fmt.Printf("Found!\n")
	}else{
		fmt.Printf("Not found!\n")
	}

}
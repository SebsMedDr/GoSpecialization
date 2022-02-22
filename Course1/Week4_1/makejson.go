package main

import "fmt"
import "encoding/json"
import "bufio"
import "os"

func main(){
	address_book := make(map[string]string)
	var name string
	var address string

	fmt.Printf("Name and address of an indivual as a space separated string. \n")
	fmt.Printf("Enter a name: \n")
	name_scanner := bufio.NewScanner(os.Stdin)
	name_scanner.Scan()
	name = name_scanner.Text()
	fmt.Printf("Enter an address: \n")
	address_scanner := bufio.NewScanner(os.Stdin)
	address_scanner.Scan()
	address = address_scanner.Text()

	address_book["name"] = name
	address_book["address"] = address
	address_book_json, err := json.Marshal(address_book)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(address_book_json))
	fmt.Printf("\n\n")
	
}
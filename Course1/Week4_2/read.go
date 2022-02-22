package main

import "fmt"
import "io/ioutil"
import s "strings"

type name struct{
	fname string
	lname string
}

func main(){

	fmt.Printf("Enter the location and name of a file: \n")
	var file_name string
	fmt.Scan(&file_name)

	content, err := ioutil.ReadFile(file_name)

	if err != nil{
		fmt.Println(err)
	}

	var name_str string
	names := []name{}
	for _,val := range content{
		if string(val) == "\n"{
			name_slice := s.Split(name_str," ")
			var name_struct name
			name_struct.fname = name_slice[0]
			name_struct.lname = s.TrimSpace(name_slice[1])
			names = append(names,name_struct)
			name_str = ""
			continue
		}
		name_str += string(val)
	}
	name_slice := s.Split(name_str," ")
	var name_struct name
	name_struct.fname = name_slice[0]
	name_struct.lname = name_slice[1]
	names = append(names,name_struct)

	for _,name_struct := range names{
		fmt.Printf("%s %s\n\n",name_struct.fname,name_struct.lname)
	}

}
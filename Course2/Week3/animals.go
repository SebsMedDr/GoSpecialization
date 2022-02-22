package main

import "fmt"
import s "strings"
import "os"

type Animal struct{
	food string
	locomotion string
	noise string
}

func (animal *Animal) Eat(){
	fmt.Printf(animal.food)
}

func (animal *Animal) Move(){
	fmt.Printf(animal.locomotion)
}

func (animal *Animal) Speak(){
	fmt.Printf(animal.noise)
}

func Animal_map() map[string](*Animal){
	cow := &Animal{food:"grass",locomotion:"walk",noise:"moo"}
	bird := &Animal{food:"worms",locomotion:"fly",noise:"peep"}
	snake := &Animal{food:"mice",locomotion:"slither",noise:"hsss"}
	return map[string]*Animal{"cow":cow,
		"bird":bird,
		"snake":snake,
	}
}

func User_animal(animal_map map[string](*Animal)){
	fmt.Printf("Please select an animal from the following list: [cow, bird, snake]\n")
	var animal_str string
	_,err := fmt.Scan(&animal_str)

	if err == nil{
		if animal_str == "cow" || animal_str == "bird" || animal_str == "snake"{
			User_animal_info(animal_map[animal_str],animal_str)
		}else if s.ToLower(animal_str) == "x"{
			os.Exit(0)
		}else{
			fmt.Printf("Not a possible input for an animal. Please try again.\n")
		}
	}else{
		fmt.Printf("Not a possible input for an animal. Please try again.\n")
	}
}

func User_animal_info(animal *Animal, animal_str string){
	fmt.Printf("What would you like to know about %s?\n",animal_str)
	fmt.Printf("You can know input: [eat,move,speak]\n")
	var animal_info string
	_,err := fmt.Scan(&animal_info)

	if err == nil{
		if s.ToLower(animal_info) == "x"{
			os.Exit(0)
		}
		Print_animal_info(animal,animal_info,animal_str)
	}else{
		fmt.Printf("Not a possible input for an animal information. Please try again.\n")
	}
}

func Print_animal_info(animal *Animal, animal_info, animal_str string){
	for {
		if animal_info == "eat"{
			fmt.Printf("The %s eats ",animal_str)
			animal.Eat()
			fmt.Print(".\n")
		}else if animal_info == "move"{
			fmt.Printf("The %s moves by ",animal_str)
			animal.Move()
			fmt.Print(".\n")
		}else if animal_info == "speak"{
			fmt.Printf("The %s sounds like ",animal_str)
			animal.Speak()
			fmt.Print(".\n")
		}else{
			fmt.Printf("Not a possible input for an animal information. Please try again.\n")
			fmt.Printf("You can know input: [eat,move,speak]\n")
			_,err := fmt.Scan(&animal_info)
			if err == nil{
				continue
			}else{
				fmt.Printf("Not a possible input for an animal information. Please try again.\n")
				fmt.Printf("You can know input: [eat,move,speak]\n")
			}
		}
		break
	}
}

func main(){
	
	fmt.Printf("Building animal map now.\n")
	animal_map := Animal_map()
	fmt.Printf("Input X at any point to exit the program.\n")
	for {
		User_animal(animal_map)
	}
}
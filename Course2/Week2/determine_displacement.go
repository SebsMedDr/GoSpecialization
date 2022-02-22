package main

import "fmt"

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64{
	return func(t float64) float64 {
		return 0.5*a*t*t+v0*t+s0
	}
}

func main(){

	var a float64
	var v0 float64
	var s0 float64
	var t float64

	fmt.Printf("Please enter an acceleration: \n")
	fmt.Scan(&a)

	fmt.Printf("Please enter an initial velocity: \n")
	fmt.Scan(&v0)

	fmt.Printf("Please enter an initial position: \n")
	fmt.Scan(&s0)

	fmt.Printf("Please enter a time: \n")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a,v0,s0)

	fmt.Printf("For an acceleration of %.2f, an initial velocity of %.2f, an initial position of %.2f, and a time of %.2f the final position is: \n",a,v0,s0,t)
	fmt.Println(fn(t))

}
package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main(){
	fmt.Println("Random Number:", rand.Intn(100))
	fmt.Printf("Square root of 5 = %g \n", math.Sqrt(5))
	fmt.Printf("Pi = %g \n", math.Pi)
}
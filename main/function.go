package main

import "fmt"

func main(){
	fmt.Println(add(2, 132))
	fmt.Println(add(2, 132))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}


// Basic function call

func add(a int, b int) int {
	return a + b
}

// we can emit the type declarations like the following

func anotherAdd (a, b int) int{
	return a + b
}

// multiple result function

func swap(a, b string) (string, string){
	return b, a
}
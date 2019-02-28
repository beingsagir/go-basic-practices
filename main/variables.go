package main

import "fmt"

var c, python, java bool = true, true, false

func main() {
	assignmentToVar()
	var i int
	fmt.Println(c, python, java, i)

}

// default value

func assignmentToVar() {
	var java int = 112
	fmt.Println(java)
}

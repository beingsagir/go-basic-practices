package main

import (
	"fmt"
	"strconv"
)


func main(){

	var a [10]int
	fmt.Println(a)
	a[2] = 667
	fmt.Println(a)

	a[9] = 22
	fmt.Println("a's length:", len(a))

	var twoDarray [3][3]string
	for i := 0; i<3;i++{
		for j:= 0; j<3; j++{
			twoDarray[i][j] = prepareString(i,  j)
		}
	}
	fmt.Println(twoDarray)

}


func prepareString(i int, j int) string{
	stringA := strconv.Itoa(i) + "_" +strconv.Itoa(j)
	return stringA
}

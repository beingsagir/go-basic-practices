package main

import (
	"fmt" 
	"strconv"
)

func main(){
	//First question: how to get int string?

    intValue := 123
    // keeping it in sheparate variable cause 
    strValue := strconv.Itoa(intValue) 
    fmt.Println(strValue)

//Second question: how to concat two string?

    firstStr := "ab"
    secondStr := "c"
    s := firstStr + secondStr
    fmt.Println(s)
}
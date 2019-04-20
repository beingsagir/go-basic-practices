package main


import "fmt"


/* Range
	Range is for for each. It helps looping the loopable contents for example slices arrays or maps
*/

func main(){
	numbers := []int{4, 1, 5}
	sum := 0

	for index, number := range numbers{
		//  Here we have got each number from array
		fmt.Println("number: (index:", index, ")=", number)
		sum += number
	}
	fmt.Println("sum", sum)


	keyValues := map[string]string{"one": "Iphone", "two": "One Plus"}
	// Loop through only keys of a map
	for key:= range keyValues{
		fmt.Println("key:", key)
	}

	// Question : How to get key index of a map?
	//------------------------------->>

	// Loop through key values of a map
	for key, value := range keyValues{
		fmt.Printf("%s -> %s \n", key, value)
	}
}




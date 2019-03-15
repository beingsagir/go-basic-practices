package main

import "fmt"

func main() {

	k := make([]int, 3)
	fmt.Println(k)

	k[0] = 20
	k[2] = 35
	k[1] = 39

	fmt.Println(k)

	k = append(k, 29, 78, 37)
	fmt.Println("After Appending", k)

	fmt.Println(k[5])

	n := k[3:]
	fmt.Println(n)

	m := k[:3]
	fmt.Println(m)

	

}

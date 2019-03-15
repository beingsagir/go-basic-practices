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

	p := []string{"Sagir", "Aziz", "Atanu", "Anish"}
	fmt.Println(p)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}

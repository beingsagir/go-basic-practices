package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["first"] = 001
	m["second"] = 002

	fmt.Println(m)
	fmt.Println(len(m))

	delete(m, "first")
	fmt.Println(m)

	_, prs := m["second"]
	fmt.Println("prs", prs)

	data, prs := m["second"]
	fmt.Println("prs", prs)
	fmt.Println("data", data)

}

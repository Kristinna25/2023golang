package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 0, 0} //slice literal
	s[4] = 99
	for _, value := range s {
		fmt.Println(value)
	}
}

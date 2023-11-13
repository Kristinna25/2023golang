package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 6)
	for _, value := range s {
		fmt.Println(value)
	}
}

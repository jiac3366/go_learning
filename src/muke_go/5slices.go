package main

import "fmt"

var capacity int = 8

func main() {
	s1 := make([]int, 0, capacity)
	s2 := append(s1, 4)
	fmt.Printf("s1 %v, len %d, cap %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s1 %v, len %d, cap %d\n", s2, len(s2), cap(s2))
	//println(s2)
}

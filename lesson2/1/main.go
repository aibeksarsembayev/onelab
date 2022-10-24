package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // cap 10
	s1 := slice[2:5]                             // 2 3 4
	// s2 := slice[2:6:7] // 2 3 4 5

	// s2 = append(s2, 100) //
	// fmt.Println("s2 100", cap(s2))
	// s2 = append(s2, 200)
	// fmt.Println("s2 200", cap(s2))

	s1[2] = 20
	fmt.Println(s1, cap(s1))

	// fmt.Println(s1, cap(s1), s2, cap(s2), slice, cap(slice))
}

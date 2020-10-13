package main

import "fmt"

func main() {
	var s []int // zero value for slice is nil
	for i := 0; i < 100; i++ {
		s = append(s, 2 * i)
	}

	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	fmt.Println(s1)

	// make a 16 length, 32 capacity slice
	s2 := make([]int, 16, 32)
	fmt.Println(s2)

	copy(s2, s1)
	fmt.Println(s2)

	// delete some elements
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)

	font := s2[0]
	s2 = s2[1:]

	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]

	fmt.Println(font, tail, s2)
}

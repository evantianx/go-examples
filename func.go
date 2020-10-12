package main

import "fmt"

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func add(a, b int) int {
	return a + b
}
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
func main()  {
	fmt.Println(div(5, 4))
	fmt.Println(div2(5, 4))
	fmt.Println(apply(add, 5, 2))
	fmt.Println(sum(2,3,4,5,6))
}

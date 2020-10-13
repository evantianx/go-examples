package main

import "fmt"

func printArray(arr [5]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}


func main()  {
	var arr1 [4]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 3, 4,12,1,45}
	var grid [3][4]int

	fmt.Println(arr1, arr2, arr3, grid)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for _, v := range arr2 {
		fmt.Println(v)
	}
}


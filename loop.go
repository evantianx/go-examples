package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

func main()  {
	fmt.Println(convertToBin(5), convertToBin(13))
}
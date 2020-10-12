package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	r := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		r = "F"
	case score < 80:
		r = "E"
	case score < 90:
		r = "B"
	case score <= 100:
		r = "A"
	}

	return r
}

func main() {
	const filename = "abc.txt"
	// we can declare variables inside if block
	// note that variables will only available inside if block
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(69),
		grade(100),

		)
}
package main

import "fmt"

func enums() {
	const (
		cpp = iota
		java
		py
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)

	fmt.Println(cpp, java, py, golang)
	fmt.Println(b, kb, mb, gb)
}

func main() {
	enums()
}
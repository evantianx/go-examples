package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "evan",
		"site": "evantian.me",
	}

	m1 := make(map[string]int)

	var m2 map[string]int

	fmt.Println(m, m1, m2)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Printf(m["site"])
	fmt.Printf(m["sate"]) // print empty

	if v, ok := m["age"]; ok {
		fmt.Printf(v)
	} else {
		fmt.Printf("key doesn't exist")
	}
 
	delete(m, "name")
}

package main

import "fmt"

func slap() {
	var a, b = 1, 3
	fmt.Println(a)
	fmt.Println(b)
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

package main

import "fmt"

func main() {

	fmt.Println("hello,this is my first program written in golang")
	fmt.Printf("Format print test %s %d\n", "this is a string", 123)

	fmt.Println("3 + 2 = ", add(3, 2))
	fmt.Println("this")

	

}

func add(a, b int) int {

	return a + b

}

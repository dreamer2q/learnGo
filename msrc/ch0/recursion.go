package main

import "fmt"

func recursion(){
	recursion()
}

func main(){

	recursion()

	fmt.Println("Recursion")

}

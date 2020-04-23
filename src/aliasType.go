package main

import "fmt"

type myType struct {
	msg string
}

type aliasType = myType

func main() {
	var ty = &aliasType{
		msg: "this is an alias type",
	}

	fmt.Printf("%+v\n", ty)
}

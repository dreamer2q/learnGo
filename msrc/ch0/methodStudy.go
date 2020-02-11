package main

import(
	"fmt"
)

type Circle struct{
	radius float64
}

func (c Circle) getArea() float64{
	return c.radius*c.radius*3.14
}

func main(){

	var c1 Circle
	c1.radius = 10.0

	fmt.Println("The area is",c1.getArea())
		
	fmt.Println("Study method of go.")
}

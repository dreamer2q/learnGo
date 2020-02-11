package main

import(
	"fmt"
	"reflect"
)


func main(){
	
	var x float64 = 3.4

	p := reflect.ValueOf(&x)

	fmt.Println("type of p:",p.Type())

	fmt.Println("p Canset ",p.CanSet())

	v := p.Elem()

	fmt.Println("v Canset ",v.CanSet())

	v.SetFloat(12.33)

	fmt.Println("v :",v.Interface())
	fmt.Println(v)



}

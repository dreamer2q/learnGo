package main

import "fmt"

type Phone interface{
	call()
}

type NokiePhone struct{

}
type IPhone struct{

}

func (nokie NokiePhone)call(){
	fmt.Println("I am NokiePhone, I can call you")
}

func (iPhone IPhone) call(){
	fmt.Println("I am IPhone, I can call you")
}

func main(){

	var phone Phone;
	phone = new(NokiePhone)
	phone.call()
	phone = new(IPhone)
	phone.call()


	
}

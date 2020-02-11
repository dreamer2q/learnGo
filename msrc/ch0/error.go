package main

import "fmt"
//import "errors"

type Err struct{
	err string
}

func (err Err)New(e string)(Err){
	err.err = e
	return err
}

func (err Err)Error()string{
	return err.err
}



func sqrt(f float64) (float64,error){

	if f<0 {
		var  err Err
		return 0,err.New("this is an error")
	}
	return 0,nil
}

func main(){

	_,err := sqrt(-1)
	if err != nil {
		fmt.Println("Error:",err.Error())
	
	}
}

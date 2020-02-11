package main

import "fmt"

func main(){

	var arr [100]int

	for i:=0;i<100; i++ {
		arr[i] = i;
	}


	for index,ele := range(arr) {
	
		fmt.Println("Index = ",index,"element = ",ele)

	}

}

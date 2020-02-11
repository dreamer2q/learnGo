 package main

 import "fmt"

 func main(){
	
	var slice1 []int = make([]int,10,20)

	var slice2 []int
 
	slice2 = []int{1,2,3,4,5,6,7,8,9}

	for i,s := range slice2 {
		fmt.Println("Index =",i,"content =",s)
	}

	slice3 := slice1[1:9]
	for i,_ := range slice3{
		slice3[i] = i
	}

	for i:=0;i<15;i++{
		slice1 = append(slice1,i*2)
	}

	fmt.Println("len = ",len(slice1),"cap =",cap(slice1))
	fmt.Println("Slice =",slice1)
 
 }

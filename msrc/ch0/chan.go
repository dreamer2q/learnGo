package main

import "fmt"
import "time"

func say(s string){

	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}

}

func sum(s []int,c chan int){
	sum := 0
	for _,v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int,c chan int){
	x,y := 0,1
	for i:=0 ;i<n;i++ {
		c <- x
		x,y = y,x+y
	}
	close(c)

}


func main(){

	c := make(chan int,60)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

}

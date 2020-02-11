package main

import "fmt"

func main(){
	
	map1 := map[string]string{
		"test":"this is the content of key 1",
		"second":"Content of second",
		"Map":"the powerful map tool you should deserve"}

	for key,value := range map1{
		fmt.Println("key=",key,"value=",value)
	}

	map1["add"] = "add a new element"
	fmt.Println(map1)
	map1["add"] = "revise key add content"
	fmt.Println(map1)



	delete(map1,"add")
	fmt.Println(map1)


	


}

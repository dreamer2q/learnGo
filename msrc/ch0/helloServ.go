package main

import (
	"io"
	"log"
	"net/http"
)



func helloHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w, "Hello. world!")
}

func main(){

	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":5555",nil)

	if err != nil {
		log.Fatal("ListenServe: ", err.Error())
	
	}
}


/* package implements 

a simple server

*/
package main

import (
//	"log"
	"net/http"
)

//definition const
const (
	DIR = "/"
)


//Main func start a file server 
func main(){

	h := http.FileServer(http.Dir(DIR)) 
	http.ListenAndServe(":999", h)


}

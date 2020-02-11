package main

import(
	"fmt"
	"net/http"
	"os"
	"io"
)

func main(){
	
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage %s :url\n",os.Args[0])
		os.Exit(1)
	}

	
	url := os.Args[1]

	resp, err := http.Get(url)
	checkErr(err)

	defer resp.Body.Close()

	io.Copy(os.Stdout,resp.Body)


	os.Exit(0)
}

func checkErr(err error){
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error:%s\n", err.Error())
	}
}

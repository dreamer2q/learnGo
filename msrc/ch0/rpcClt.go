package main

import(
	"net/rpc"
	"log"
	"fmt"
)

type Args struct {
	A,B int
}
type Quotient struct{
	Quo,Rem int
}

func main(){


	client,err := rpc.DialHTTP("tcp","localhost:1234")
	if err != nil {
		log.Fatal("RPC DialHTTP Error:",err.Error())
		return
	}

	args := &Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	if err != nil{
		log.Fatal("Arith error:",err.Error())
	}

	fmt.Printf("Arith: %d*%d = %d\n",args.A,args.B,reply)

	quotient := new(Quotient)
	
	client.Call("Arith.Divide", args, &quotient)
	if err != nil{
		log.Fatal("Arith.Divide error:",err.Error())	
    }
	fmt.Printf("Arith.Divide: %d/%d = %d ... %d\n", args.A,args.B,quotient.Quo,quotient.Rem)
	
	args.B = 0;
	err = client.Call("Arith.Divide", args, &quotient)
	if err != nil{
		log.Fatal("Arith.Divice error:",err.Error())	
	}
	fmt.Printf("Arith.Divide: %d/%d = %d ... %d\n", args.A,args.B,quotient.Quo,quotient.Rem)

}

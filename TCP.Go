package main

import (
	"fmt"
)


func main(){
	comms := make(chan string,1)

	go server(comms)
	go client(comms)
	

	for {

	}
}

func client(ch chan string){
ch <- "syn seq"
fmt.Println("c1")
msg := <- ch
fmt.Println("c2")
if(msg != "syn seq ACK"){
//Something wrong (repeat send)
fmt.Println("c-wrong")
}else{
fmt.Println("c-right")
ch <- msg + " DATA"
}
}

func server(ch chan string){
msg := <- ch
ch <- msg + " ACK"
data := <- ch
fmt.Println(data)
}
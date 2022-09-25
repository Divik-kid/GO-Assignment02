package main

import (
	"fmt"
	"time"
)

func main() {
	comms := make(chan packet, 1)

	go server(comms)
	go client(comms)

	for {

	}
}

// considering if ordernumber matters
func client(ch chan packet) {
	seq := 0
	ack := 0
	for {
		//Step 1 send x = syn seq'
		time.Sleep(1 * time.Second)
		ch <- packet{"", seq, ack, 0} //
		fmt.Println("sending step1")
		//Step 2 recieve x+1 and y
		synAck := <-ch
		if synAck.acknowledgement == seq+1 && synAck.sequence == ack+100 {
			fmt.Println("sending data")
			//Step 3 send y+1 and x+1 and data
			ch <- packet{"Here is my data", synAck.acknowledgement, synAck.sequence, synAck.orderNumber + 1}
			seq++
			ack = ack + 100
		}
	}
}

func server(ch chan packet) {
	exSec := 0
	exAck := 0

	for {

		//Step 1 listen for requests, check x = syn
		syn := <-ch
		fmt.Println("recieving step1")
		if syn.sequence == exSec {
			//Step 2 send, x+1 and y
			ch <- packet{"", syn.acknowledgement + 100, syn.sequence + 1, 0}
			fmt.Println("sending step2")
		}
		//Step 3 listen for requests, check x+1 = syn and y+1 = ack

		if syn.acknowledgement == exAck+100 && syn.sequence == exSec+1 {
			fmt.Println("recieving data")
			//Step 4 listen for data
			exSec++
			exAck = exAck + 100
			fmt.Print(syn.data + " number: ")
			fmt.Println(syn.sequence)

			//lets say we need 15 packages, then we could insert the packages int an array of packages according to their sequence/package number
			//possibly even have the server request missing packages
		}
	}
}

type packet struct {
	data            string
	sequence        int
	acknowledgement int
	orderNumber     int
}

func (p packet) Marshalling() []packet {
	return nil
}

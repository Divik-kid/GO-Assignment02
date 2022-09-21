package main

import (
	"fmt"
)

func main() {
	comms := make(chan packet, 1)

	go server(comms)
	go client(comms)

	for {

	}
}

func client(ch chan packet) {
	for {
		ch <- packet{"", 1, 0, 0} //
		fmt.Println("sending step1")
		synAck := <-ch
		if synAck.acknowledgement == 2 {
			fmt.Println("sending data")
			ch <- packet{"Here is my data", synAck.acknowledgement + 1, synAck.sequence + 1, 0} //
		}
	}
}

func server(ch chan packet) {
	for {
		syn := <-ch
		fmt.Println("recieving step1")
		if syn.sequence == 1 {
			ch <- packet{"", 100, 2, 0}
			fmt.Println("sending step2")
		}
		if syn.acknowledgement == 101 {
			fmt.Println("recieving data")
			data := <-ch
			fmt.Println(data.data)
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

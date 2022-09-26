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

	dataToSend := []string{"here", "is", "data"}

	for index, element := range dataToSend {
		//Step 1 send x = syn seq'
		time.Sleep(1 * time.Second)
		ch <- packet{"", seq, ack, 0, 0}
		fmt.Println("sending step1")
		//Step 2 recieve x+1 and y
		synAck := <-ch
		if synAck.acknowledgement == seq+1 && synAck.sequence == ack+100 {
			fmt.Println("sending data")
			//Step 3 send y+1 and x+1 and data
			ch <- packet{element, synAck.acknowledgement, synAck.sequence, index, len(dataToSend)}
			seq++
			ack = ack + 100
		}
	}
}

func server(ch chan packet) {
	exSec := 0
	exAck := 0

	recivedData := make([]string, 0)

	for {

		//Step 1 listen for requests, check x = syn
		syn := <-ch
		fmt.Println("recieving step1")
		if syn.sequence == exSec {
			//Step 2 send, x+1 and y
			ch <- packet{"", syn.acknowledgement + 100, syn.sequence + 1, 0, 0}
			fmt.Println("sending step2")
		}
		//Step 3 listen for requests, check x+1 = syn and y+1 = ack

		if syn.acknowledgement == exAck+100 && syn.sequence == exSec+1 {
			fmt.Println("recieving data")
			//Step 4 listen for data
			exSec++
			exAck = exAck + 100

			if syn.sequence == 1 {
				recivedData = make([]string, syn.totalAmount)
			}

			recivedData[syn.orderNumber] = syn.data

			if syn.sequence == syn.totalAmount {
				fmt.Printf("%+q", recivedData)
				break
			}
		}
	}
}

type packet struct {
	data            string
	sequence        int
	acknowledgement int
	orderNumber     int
	totalAmount     int
}

func (p packet) Marshalling() []packet {
	return nil
}

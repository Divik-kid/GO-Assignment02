# GO Assignment 2 - TCP/IP Simulator in Go

## a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

The only packages that gets imported are fmt and time.
Hence this is a rather low level implementation, using mostly go's build in features.
The data that is transfered, is reprecented by strings.
The data and metadata is transfered using a struct, defined in go.
The struct contains data and meta-data used to facilitate the handshake.
This only works because every part of the system runs in go.

## b) Does your implementation use threads or processes? Why is it not realistic to use threads?

We use threads :)
It is not a realistic implementation, since real networks interface between hosts, not threads.
Hence your implementation only works if the two actors have acess to the same memeory.

## c) How do you handle message re-ordering?

Once we recive a packet we check the number of the packet, as defined by the client.
The data contained inside the packet is then inserted into a list, on the index specified by the client.
That way, the data is inserted into the correct order when it is recived by the server.

## d) How do you handle message loss?

We expect that the program handles message loss by requesting the missing packets again.
Specifically, if a packet is recived that should not be recived yet, the client is requested, to send packets again.
The client will resend every package starting from the first missing packet in the send order.

## e) Why is the 3-way handshake important?

It is important, in order to know that the server or recipient, is able to recive the data.
One instance where that is important, is with banking.
When making banking transfers it is really important that every host in the system knows that transfers are recived.

It is impossible to prove that the data got to the recipient, but the handshake can make it likely to happen.

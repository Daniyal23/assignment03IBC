package main

import (
	a1 "github.com/Daniyal23/assignment01IBC"
	"encoding/gob"
	"log"
	"net"
	"fmt"


)



func OwnInfo(port string, name string, conn net.Conn) {
	gobEncoder := gob.NewEncoder(conn)
	err := gobEncoder.Encode(&port)
	if err != nil {
		log.Println(err)
	}

	gobEncoder = gob.NewEncoder(conn)
	err = gobEncoder.Encode(&name)
	if err != nil {
		log.Println(err)
	}
}



func main() {
  var myaddress string
	conn, err := net.Dial("tcp", "localhost:6000")
	if err != nil {

		//handle error

	}
	myaddress=conn.LocalAddr().String()
	var recvdBlock *a1.Block
	dec := gob.NewDecoder(conn)
	err = dec.Decode(&recvdBlock)
	if err != nil {

		//handle error

	}
	fmt.Println(recvdBlock)

	var arrayRec []string
	dec = gob.NewDecoder(conn)
	err = dec.Decode(&arrayRec)
	if err != nil {

	fmt.Println(err)

	}
	fmt.Println(arrayRec)

	n, err := net.Listen("tcp", myaddress)
	if err != nil {
		fmt.Println("pfft")
		log.Println(err)
	}


		fmt.Println("pfft")
		conn, err = n.Accept()
		if err != nil {
			log.Println(err)
		}

		var recvdBlock1 a1.Block
		dec = gob.NewDecoder(conn)
		err = dec.Decode(&recvdBlock1)
		if err != nil {

			//handle error

		}
		fmt.Println(recvdBlock1)
		var coins int
		for recvdBlock.PreviousBlock != nil {
			if recvdBlock.Transaction == "Satoshi" {
				coins= coins + recvdBlock.Coins
			}

			recvdBlock = recvdBlock.PreviousBlock

		}
	 if coins>=recvdBlock1.Coins {
		 recvdBlock= a1.InsertBlock(recvdBlock1.Transaction, recvdBlock,recvdBlock1.Coins)

	 }

	 fmt.Println("The coins are ", coins)


}

package main

import (
	a1 "github.com/Daniyal23/assignment01IBC"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"
)

func SendBlockchain(port net.Conn, chainHead *a1.Block) {
	gobEncoder := gob.NewEncoder(port)
	err := gobEncoder.Encode(&chainHead)
	if err != nil {
		log.Println(err)
	}
}

func SendPeers(port net.Conn, peers []string) {
	gobEncoder := gob.NewEncoder(port)
	err := gobEncoder.Encode(peers)
	if err != nil {
		log.Println(err)
	}
}

func handleConnection(c net.Conn, chainHead *a1.Block) {
	chainHead = a1.InsertBlock("Satoshi", chainHead, 100)
}

func main() {
	var chainHead *a1.Block
	chainHead = a1.InsertBlock("GenesisBlock", nil,100)
	fmt.Println("is this working ?")
	var conSlice []net.Conn

	n, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println("pfft")
		log.Println(err)
	}

	for i:=0;i<2;i++ {
		fmt.Println("pfft")
		conn, err := n.Accept()
		if err != nil {
			log.Println(err)
		}
		conSlice = append(conSlice, conn)
		println("\nGot connected")
		chainHead = a1.InsertBlock("Satoshi", chainHead, 100)
	}
	var array []string
	for i:=0;i<len(conSlice);i++{

		array=append(array, conSlice[i].RemoteAddr().String())
	}

	fmt.Println(array)
	for i:=0;i<len(conSlice);i++{
		SendBlockchain(conSlice[i], chainHead)
		go SendPeers(conSlice[i], array)
		time.Sleep(3*time.Second)
	}

  var chainHead2 a1.Block
	chainHead2.Transaction="Alice to Bob"
	chainHead2.Coins= 100
	conn, err := net.Dial("tcp", array[0])
	if err != nil {

		//handle error

	}
	gobEncoder := gob.NewEncoder(conn)
	err = gobEncoder.Encode(chainHead2)
	if err != nil {
		log.Println(err,"Hello this is the life ")
	}



}

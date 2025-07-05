package main

import (
	"fmt"
	"log"

	"github.com/ahemtan/distFS/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	return nil
}

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListnerAddr:   ":8080",
		Decoder:       p2p.DefaultDecoder{},
		HandShakeFunc: p2p.NOPHandShakeFunc,
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("@TCP message recived %+v \n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}

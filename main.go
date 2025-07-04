package main

import (
	"log"

	"github.com/ahemtan/distFS/p2p"
)

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListnerAddr:   ":8080",
		Decoder:       p2p.DefaultDecoder{},
		HandShakeFunc: p2p.NOPHandShakeFunc,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}

package p2p

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outBound bool
}

func NewRCPPeer(conn net.Conn, outBound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outBound: false,
	}
}

type TCPTransportOpts struct {
	ListnerAddr   string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listner net.Listener

	mu   sync.RWMutex
	peer map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listner, err = net.Listen("tcp", t.ListnerAddr)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listner.Accept()
		if err != nil {
			fmt.Printf("@TCP accept error: %s\n", err)
		}

		fmt.Printf("@TCP new connection from %s\n", conn.RemoteAddr().String())

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewRCPPeer(conn, true)

	if err := t.HandShakeFunc(peer); err != nil {
		fmt.Printf("@TCP handshake error")
		conn.Close()
		return
	}

	msg := &Message{}

	for {
		if err := t.Decoder.Decoder(conn, msg); err != nil {
			fmt.Printf("@TCP decode error: %s\n", err)
			if err == io.EOF || strings.Contains(err.Error(), "connection was forcibly closed") {
				conn.Close()
				return
			}
			continue
		}

		msg.From = conn.RemoteAddr()

		fmt.Printf("@TCP message: %+v\n", msg)
	}

}

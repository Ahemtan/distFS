package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {
	ListenAddr := ":8080"
	tr := NewTCPTransport(TCPTransportOpts{ListnerAddr: ListenAddr})
	assert.Equal(t, tr.ListnerAddr, ListenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}

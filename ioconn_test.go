package ioconn

import (
	"io"
	"net"
	"testing"
)

var (
	_ io.Closer = CloserFunc(nil)
	_ net.Addr  = FileAddr("")
	_ net.Addr  = Addr{}
	_ net.Conn  = Conn{}
)

func TestT(*testing.T) {
	// Added to stop go test complaining about no tests
}

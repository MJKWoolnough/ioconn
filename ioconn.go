// Package ioconn allows any combination of an io.Reader, io.Writer and io.Closer to become a net.Conn.
package ioconn // import "vimagination.zapto.org/ioconn"

import (
	"errors"
	"io"
	"net"
	"time"
)

// CloserFunc is a func that implements the io.Closer interface allowing a
// closure or other function to be io.Closer.
type CloserFunc func() error

// Close simply calls the CloserFunc func.
func (c CloserFunc) Close() error {
	return c()
}

// FileAddr is a net.Addr that represents a file. Should be a full path.
type FileAddr string

// Network always returns "file".
func (f FileAddr) Network() string {
	return "file"
}

// String returns file://path.
func (f FileAddr) String() string {
	return "file://" + string(f)
}

// Addr is a simple implementation of the net.Addr interface.
type Addr struct {
	Net, Str string
}

// Network returns the Net string.
func (a Addr) Network() string {
	return a.Net
}

// String returns the Str string.
func (a Addr) String() string {
	return a.Str
}

// Conn implements a net.Conn.
type Conn struct {
	io.Reader
	io.Writer
	io.Closer
	Local, Remote               net.Addr
	ReadDeadline, WriteDeadline time.Time
}

// Read implements the io.Reader interface.
func (c *Conn) Read(p []byte) (int, error) {
	if !c.ReadDeadline.IsZero() && time.Now().After(c.ReadDeadline) {
		return 0, ErrTimeout
	}

	return c.Reader.Read(p)
}

// Write implements the io.Writer interface.
func (c *Conn) Write(p []byte) (int, error) {
	if !c.ReadDeadline.IsZero() && time.Now().After(c.WriteDeadline) {
		return 0, ErrTimeout
	}

	return c.Writer.Write(p)
}

// LocalAddr returns the Local Address.
func (c *Conn) LocalAddr() net.Addr {
	return c.Local
}

// RemoteAddr returns the Remote Address.
func (c *Conn) RemoteAddr() net.Addr {
	return c.Remote
}

// SetDeadline implements the Conn SetDeadline method.
func (c *Conn) SetDeadline(t time.Time) error {
	err := c.SetReadDeadline(t)
	err2 := c.SetWriteDeadline(t)

	if err != nil {
		return err
	}

	return err2
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *Conn) SetReadDeadline(t time.Time) error {
	c.ReadDeadline = t

	if rd, ok := c.Writer.(interface {
		SetReadDeadline(time.Time) error
	}); ok {
		return rd.SetReadDeadline(t)
	}

	return nil
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *Conn) SetWriteDeadline(t time.Time) error {
	c.WriteDeadline = t

	if wd, ok := c.Writer.(interface {
		SetWriteDeadline(time.Time) error
	}); ok {
		return wd.SetWriteDeadline(t)
	}

	return nil
}

// Errors.
var (
	ErrTimeout = errors.New("timeout occurred")
)

# ioconn
--
    import "github.com/MJKWoolnough/ioconn"

Package ioconn allows any combination of an io.Reader, io.Writer and io.Closer to become a net.Conn

## Usage

```go
var ErrUnimplemented = errors.New("not implmented")
```
Errors

#### type Addr

```go
type Addr struct {
	Net, Str string
}
```

Addr is a simple implementation of the net.Addr interface

#### func (Addr) Network

```go
func (a Addr) Network() string
```
Network returns the Net string

#### func (Addr) String

```go
func (a Addr) String() string
```
String returns the Str string

#### type CloserFunc

```go
type CloserFunc func() error
```

CloserFunc is a func that implements the io.Closer interface allowing a closure
or other function to be io.Closer

#### func (CloserFunc) Close

```go
func (c CloserFunc) Close() error
```
Close simply calls the CloserFunc func

#### type Conn

```go
type Conn struct {
	io.Reader
	io.Writer
	io.Closer
	Local, Remote net.Addr
}
```

Conn implements a net.Conn

#### func (Conn) LocalAddr

```go
func (c Conn) LocalAddr() net.Addr
```
LocalAddr returns the Local Address

#### func (Conn) RemoteAddr

```go
func (c Conn) RemoteAddr() net.Addr
```
RemoteAddr returns the Remote Address

#### func (Conn) SetDeadline

```go
func (Conn) SetDeadline(time.Time) error
```
SetDeadline is unimplemented and always returns an error

#### func (Conn) SetReadDeadline

```go
func (Conn) SetReadDeadline(time.Time) error
```
SetReadDeadline is unimplemented and always returns an error

#### func (Conn) SetWriteDeadline

```go
func (Conn) SetWriteDeadline(time.Time) error
```
SetWriteDeadline is unimplemented and always returns an error

#### type FileAddr

```go
type FileAddr string
```

FileAddr is a net.Addr that represents a file. Should be a full path

#### func (FileAddr) Network

```go
func (f FileAddr) Network() string
```
Network always returns "file"

#### func (FileAddr) String

```go
func (f FileAddr) String() string
```
String returns file://path

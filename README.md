# ioconn
--
    import "github.com/MJKWoolnough/ioconn"

Package ioconn allows any combination of an io.Reader, io.Writer and io.Closer
to become a net.Conn

## Usage

```go
var (
	ErrTimeout = errors.New("timeout occurred")
)
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
	Local, Remote               net.Addr
	ReadDeadline, WriteDeadline time.Time
}
```

Conn implements a net.Conn

#### func (*Conn) LocalAddr

```go
func (c *Conn) LocalAddr() net.Addr
```
LocalAddr returns the Local Address

#### func (*Conn) Read

```go
func (c *Conn) Read(p []byte) (int, error)
```
Read implements the io.Reader interface

#### func (*Conn) RemoteAddr

```go
func (c *Conn) RemoteAddr() net.Addr
```
RemoteAddr returns the Remote Address

#### func (*Conn) SetDeadline

```go
func (c *Conn) SetDeadline(t time.Time) error
```
SetDeadline implements the Conn SetDeadline method

#### func (*Conn) SetReadDeadline

```go
func (c *Conn) SetReadDeadline(t time.Time) error
```
SetReadDeadline implements the Conn SetReadDeadline method

#### func (*Conn) SetWriteDeadline

```go
func (c *Conn) SetWriteDeadline(t time.Time) error
```
SetWriteDeadline implements the Conn SetWriteDeadline method

#### func (*Conn) Write

```go
func (c *Conn) Write(p []byte) (int, error)
```
Write implements the io.Writer interface

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

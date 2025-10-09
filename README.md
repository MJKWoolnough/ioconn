# ioconn

[![CI](https://github.com/MJKWoolnough/ioconn/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/ioconn/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/ioconn.svg)](https://pkg.go.dev/vimagination.zapto.org/ioconn)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/ioconn)](https://goreportcard.com/report/vimagination.zapto.org/ioconn)

--
    import "vimagination.zapto.org/ioconn"

Package ioconn allows any combination of an io.Reader, io.Writer and io.Closer to become a net.Conn.

## Highlights

 - Wrap an `io.Reader, `io.Writer`, and `io.Closer` into a `net.Conn`.
 - Provides simple `net.Addr` implementations.

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/ioconn

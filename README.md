# Middleware [![GoDoc](https://godoc.org/github.com/itsmontoya/middleware?status.svg)](https://godoc.org/github.com/itsmontoya/middleware) ![Status](https://img.shields.io/badge/status-alpha-red.svg)
Middleware is a middleware assistant for I/O operations

## Usage
```go
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/itsmontoya/middleware"
)

func main() {
	var (
		rdr io.Reader
		val string
		err error

		mws = middleware.NewMWs(middleware.GZipMW{})
	)

	if rdr, err = write(mws, "Hello"); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if val, err = read(mws, rdr); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(val)
}

func write(mws *middleware.MWs, val string) (out io.Reader, err error) {
	var (
		wtr *middleware.Writer
		buf = bytes.NewBuffer(nil)
	)

	if wtr, err = mws.Writer(buf); err != nil {
		err = errors.New("Error getting middleware Writer:" + err.Error())
		return
	}

	if _, err = wtr.Write([]byte(val)); err != nil {
		err = errors.New("Error writing to middleware Writer:" + err.Error())
		return
	}

	if err = wtr.Close(); err != nil {
		err = errors.New("Error closing middleware Writer:" + err.Error())
		return
	}

	out = buf
	return
}

func read(mws *middleware.MWs, in io.Reader) (val string, err error) {
	var (
		rdr *middleware.Reader
		buf = bytes.NewBuffer(nil)
	)

	if rdr, err = mws.Reader(in); err != nil {
		err = errors.New("Error getting middleware Reader:" + err.Error())
		return
	}

	if _, err = io.Copy(buf, rdr); err != nil {
		err = errors.New("Error reading from middleware Reader:" + err.Error())
		return
	}

	val = buf.String()
	return
}

```
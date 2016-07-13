package middleware

import "io"

// Middleware is the interface that defines an encoder/decoder chain.
type Middleware interface {
	Name() string
	Writer(w io.Writer) (io.WriteCloser, error)
	Reader(r io.Reader) (io.ReadCloser, error)
}

// NewMWs returns a new MWs
func NewMWs(mws ...Middleware) *MWs {
	reverseMWSlice(mws)
	return &MWs{mws}
}

// MWs manages middlewares
type MWs struct {
	s []Middleware
}

// Writer returns a new middleware writer
func (m *MWs) Writer(w io.Writer) (out *Writer, err error) {
	return NewWriter(w, m.s)
}

// Reader returns a new middleware reader
func (m *MWs) Reader(r io.Reader) (out *Reader, err error) {
	return NewReader(r, m.s)
}

func reverseMWSlice(mws []Middleware) {
	var n int
	mc := len(mws) - 1
	for i := range mws {
		if n = mc - i; n == i || n < i {
			break
		}

		mws[i], mws[n] = mws[n], mws[i]
	}
}

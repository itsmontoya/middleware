package middleware

import (
	"io"

	"github.com/missionMeteora/toolkit/crypty"
)

// NewCryptyMW returns a new Crypty middleware
func NewCryptyMW(key, iv []byte) *CryptyMW {
	return &CryptyMW{key, iv}
}

// CryptyMW handles encryption
type CryptyMW struct {
	key []byte
	iv  []byte
}

// Name returns the middleware name
func (c *CryptyMW) Name() string {
	return "encryption/crypty"
}

// Writer returns a new gzip writer
func (c *CryptyMW) Writer(w io.Writer) (io.WriteCloser, error) {
	return crypty.NewWriterPair(w, c.key, c.iv)
}

// Reader returns a new gzip reader
func (c *CryptyMW) Reader(r io.Reader) (rc io.ReadCloser, err error) {
	return crypty.NewReaderPair(r, c.key, c.iv)
}

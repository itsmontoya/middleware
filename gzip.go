package middleware

import (
	"compress/gzip"
	"io"
)

// GZipMW handles gzipping
type GZipMW struct {
}

// Name returns the middleware name
func (g GZipMW) Name() string {
	return "compress/gzip"
}

// Writer returns a new gzip writer
func (g GZipMW) Writer(w io.Writer) (io.WriteCloser, error) {
	return gzip.NewWriter(w), nil
}

// Reader returns a new gzip reader
func (g GZipMW) Reader(r io.Reader) (rc io.ReadCloser, err error) {
	if rc, err = gzip.NewReader(r); err != nil {
		rc = nil
	}

	return
}

package middleware

import (
	"encoding/base64"
	"io"
	"io/ioutil"
)

// Base64Name is the name constant for the Base64 middleware
const Base64Name = "encoding/base64"

// Base64MW is a base64 middleware
// Note: This should always be the LAST of the modifying middleware
type Base64MW struct {
}

// Name returns the middleware name
func (b Base64MW) Name() string {
	return Base64Name
}

// Writer returns a new gzip writer
func (b Base64MW) Writer(w io.Writer) (io.WriteCloser, error) {
	return base64.NewEncoder(base64.StdEncoding, w), nil
}

// Reader returns a new gzip reader
func (b Base64MW) Reader(r io.Reader) (rc io.ReadCloser, err error) {
	return ioutil.NopCloser(base64.NewDecoder(base64.StdEncoding, r)), nil
}

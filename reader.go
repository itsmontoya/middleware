package middleware

import (
	"io"

	"github.com/missionMeteora/toolkit/errors"
)

// NewReader returns a new middleware reader
func NewReader(r io.Reader, mws []Middleware) (out *Reader, err error) {
	var (
		rdr io.ReadCloser
		mwl = len(mws)
	)

	out = &Reader{r: r, rcs: make([]io.ReadCloser, mwl)}

	for i, mw := range mws {
		if i == 0 {
			if rdr, err = mw.Reader(r); err != nil {
				goto END
			}
		} else {
			if rdr, err = mw.Reader(rdr); err != nil {
				goto END
			}
		}

		out.rcs[mwl-1-i] = rdr
	}

END:
	if err != nil {
		out.Close()
		out = nil
	}

	return
}

// Reader is the middleware readr interface
type Reader struct {
	r   io.Reader
	rcs []io.ReadCloser
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if len(r.rcs) > 0 {
		return r.rcs[0].Read(b)
	}

	return r.r.Read(b)
}

// Close will close this readr (and it's underlying middleware readrs)
func (r *Reader) Close() (err error) {
	var errs errors.ErrorList
	for _, wc := range r.rcs {
		if wc == nil {
			continue
		}

		errs.Push(wc.Close())
	}

	return errs.Err()
}

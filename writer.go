package middleware

import (
	"io"

	"github.com/missionMeteora/toolkit/errors"
)

// NewWriter returns a new middleware writer
func NewWriter(w io.Writer, mws []Middleware) (out *Writer, err error) {
	var (
		wtr io.WriteCloser
		mwl = len(mws)
	)

	out = &Writer{wcs: make([]io.WriteCloser, mwl)}

	for i, mw := range mws {
		if i == 0 {
			if wtr, err = mw.Writer(w); err != nil {
				goto END
			}
		} else {
			if wtr, err = mw.Writer(wtr); err != nil {
				goto END
			}
		}

		out.wcs[mwl-1-i] = wtr
	}

END:
	if err != nil {
		out.Close()
		out = nil
	}

	return
}

// Writer is the middleware writer interface
type Writer struct {
	wcs []io.WriteCloser
}

func (w *Writer) Write(b []byte) (n int, err error) {
	return w.wcs[0].Write(b)
}

// Close will close this writer (and it's underlying middleware writers)
func (w *Writer) Close() (err error) {
	var errs errors.ErrorList
	for _, wc := range w.wcs {
		if wc == nil {
			continue
		}

		errs = errs.Append(wc.Close())
	}

	return errs.Err()
}

package middleware

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
)

var (
	cryptyKey = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	cryptyIV  = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	valShort  = `Hello`
	valMedium = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`
	valLong   = `Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. 
	
	Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? 
	
	Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"`
)

func TestGZip(t *testing.T) {
	if err := testSuite(NewMWs(GZipMW{})); err != nil {
		t.Error(err)
	}
}

func TestCrypty(t *testing.T) {
	if err := testSuite(NewMWs(NewCryptyMW(cryptyKey, cryptyIV))); err != nil {
		t.Error(err)
	}
}

func TestB64(t *testing.T) {
	if err := testSuite(NewMWs(Base64MW{})); err != nil {
		t.Error(err)
	}
}

func TestAll(t *testing.T) {
	if err := testSuite(NewMWs(
		GZipMW{},
		NewCryptyMW(cryptyKey, cryptyIV),
		Base64MW{},
	)); err != nil {
		t.Error(err)
	}
}

func testSuite(mws *MWs) (err error) {
	if err = test(mws, valShort); err != nil {
		return
	}

	if err = test(mws, valMedium); err != nil {
		return
	}

	if err = test(mws, valLong); err != nil {
		return
	}

	return
}

func test(mws *MWs, val string) (err error) {
	var (
		wtr *Writer
		rdr *Reader

		buf  = bytes.NewBuffer(nil)
		nbuf = bytes.NewBuffer(nil)
	)

	if wtr, err = mws.Writer(buf); err != nil {
		return errors.New("Error getting middleware Writer:" + err.Error())
	}

	if _, err = wtr.Write([]byte(val)); err != nil {
		return errors.New("Error writing to middleware Writer:" + err.Error())
	}

	if err = wtr.Close(); err != nil {
		return errors.New("Error closing middleware Writer:" + err.Error())
	}

	if rdr, err = mws.Reader(buf); err != nil {
		return errors.New("Error getting middleware Reader:" + err.Error())
	}

	if _, err = io.Copy(nbuf, rdr); err != nil {
		return errors.New("Error reading from middleware Reader:" + err.Error())
	}

	if str := nbuf.String(); str != val {
		return fmt.Errorf("Middleware output does not match, the returned value '%s' should have been '%s'", str, val)
	}

	return
}

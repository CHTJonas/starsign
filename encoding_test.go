package starsign

import (
	"io"
	"strings"
	"testing"
)

func TestDecodeSignature(t *testing.T) {
	var err error
	_, err = DecodeSignature(makeReader("just a load of old nonsense"))
	if err != ErrMalformedSig {
		t.Error("Malformed signature test failed")
	}
	_, err = DecodeSignature(makeReader(invalidSignatureNotStarsign))
	if err != ErrNotStarsignSig {
		t.Error("Non-Starsign PEM type test failed")
	}
	_, err = DecodeSignature(makeReader(invalidSignatureHeader))
	if err != ErrUntrustHeaders {
		t.Error("Untrusted PEM headers test failed")
	}
	_, err = DecodeSignature(makeReader(invalidSignatureTooShort))
	if err != ErrSigDataTooShort {
		t.Error("Short signature decoding test failed")
	}
	_, err = DecodeSignature(makeReader(validSignature))
	if err != nil {
		t.Error("Signature decoding test failed")
	}
}

func makeReader(str string) io.Reader {
	return strings.NewReader(str)
}

var invalidSignatureNotStarsign = "-----BEGIN SOME OTHER SIGNATURE-----" + validSignature[34:222] + "-----END SOME OTHER SIGNATURE-----"
var invalidSignatureHeader = validSignature[:35] + "Header: should be rejected\n" + validSignature[34:]
var invalidSignatureTooShort = validSignature[:35] + "dGVzdA==" + validSignature[221:]

var validSignature = `-----BEGIN STARSIGN SIGNATURE-----
xOs34GI4L4ED9qGfD/oxXMlMCka227Dy9ei2m2XjkhGDXLYeYEWL9A8rCish77bC
9HvQXj77SWY7oNnAhc2IIQAAACAweG0cWt9Pm3bvuqS7hr3qCxt9jVzmxpTE8M8A
mOmS3wAAACAKaIJ1I5aRAFZkL7xM2dJ/lL5OUcxWZxhK6ilgCuHegw==
-----END STARSIGN SIGNATURE-----`

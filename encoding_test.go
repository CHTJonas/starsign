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
	_, err = DecodeSignature(makeReader(invalidSignatureTooShort1))
	if err != ErrSigDataTooShort {
		t.Error("Short signature decoding test 1 failed")
	}
	_, err = DecodeSignature(makeReader(validSignature))
	if err != nil {
		t.Error("Signature decoding test failed")
	}
}

func makeReader(str string) io.Reader {
	return strings.NewReader(str)
}

var invalidSignatureNotStarsign = "-----BEGIN SOME OTHER SIGNATURE-----" + validSignature[34:259] + "-----END SOME OTHER SIGNATURE-----"
var invalidSignatureHeader = validSignature[:35] + "Header: should be rejected\n" + validSignature[34:]
var invalidSignatureTooShort1 = validSignature[:35] + "dGVzdA==" + validSignature[258:]

var validSignature = `-----BEGIN STARSIGN SIGNATURE-----
5bkdFcZkytSKgUHIfLkNkyXPI3OmksfExlDHyM+gnhd1v7MI+z7zI/IWv5MQh+C1
bIyqPlVpLTj0zg8qOdPqsUgAAAAAAAAAAAAAIAcz+SjmZjstWGCRJuBkW2mLW4/P
dNsdCnfRzi/gi33VAAAAIGKXuYKFbHKSWj876mabN7w7uYLmbuUpriaWtlw8BdY7
ZWNkc2Etc2hhMi1uaXN0cDI1Ng==
-----END STARSIGN SIGNATURE-----`

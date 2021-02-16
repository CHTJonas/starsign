package starsign

import (
	"encoding/base64"
	"encoding/gob"
	"io"
)

func EncodeSignature(out io.Writer, sig *Signature) error {
	b64 := base64.NewEncoder(base64.StdEncoding, out)
	enc := gob.NewEncoder(b64)
	err := enc.Encode(sig)
	if err != nil {
		return err
	}
	if err = b64.Close(); err != nil {
		return err
	}
	return nil
}

func DecodeSignature(in io.Reader) (*Signature, error) {
	sig := new(Signature)
	b64 := base64.NewDecoder(base64.StdEncoding, in)
	dec := gob.NewDecoder(b64)
	err := dec.Decode(sig)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

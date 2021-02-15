package starsign

import (
	"encoding/base64"
	"encoding/gob"
	"io"

	"golang.org/x/crypto/ssh"
)

func EncodeSignature(out io.Writer, sig *ssh.Signature) error {
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

func DecodeSignature(in io.Reader) (*ssh.Signature, error) {
	b64 := base64.NewDecoder(base64.StdEncoding, in)
	sig := new(ssh.Signature)
	dec := gob.NewDecoder(b64)
	err := dec.Decode(sig)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

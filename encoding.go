package starsign

import (
	"encoding/gob"
	"io"

	"golang.org/x/crypto/ssh"
)

func EncodeSignature(out io.Writer, sig *ssh.Signature) error {
	enc := gob.NewEncoder(out)
	err := enc.Encode(sig)
	if err != nil {
		return err
	}
	return nil
}

func DecodeSignature(in io.Reader) (*ssh.Signature, error) {
	sig := new(ssh.Signature)
	dec := gob.NewDecoder(in)
	err := dec.Decode(sig)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

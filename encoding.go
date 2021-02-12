package starsign

import (
	"encoding/gob"
	"io"

	"golang.org/x/crypto/ssh"
)

func EncodeSignature(buf io.Writer, sig *ssh.Signature) error {
	enc := gob.NewEncoder(buf)
	err := enc.Encode(sig)
	if err != nil {
		return err
	}
	return nil
}

func DecodeSignature(buf io.Reader) (*ssh.Signature, error) {
	sig := new(ssh.Signature)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(sig)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

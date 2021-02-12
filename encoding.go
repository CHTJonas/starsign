package starsign

import (
	"bytes"
	"encoding/gob"

	"golang.org/x/crypto/ssh"
)

func EncodeSignature(sig *ssh.Signature) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(sig)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeSignature(data []byte) (*ssh.Signature, error) {
	sig := new(ssh.Signature)
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(sig)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

package starsign

import (
	"bytes"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
)

func serialise(sig *Signature) ([]byte, error) {
	if len(sig.Hash) != 64 {
		panic("Unexpected BLAKE2 hash length")
	}
	buf := new(bytes.Buffer)
	_, err := buf.Write(sig.Hash)
	if err != nil {
		return nil, err
	}
	b := make([]byte, 8)
	length := len(sig.Sig)
	binary.LittleEndian.PutUint64(b, uint64(length))
	_, err = buf.Write(b)
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(sig.Sig)
	if err != nil {
		return nil, err
	}
	_, err = buf.WriteString(sig.Type)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func deserialise(data []byte) (*Signature, error) {
	sig := new(Signature)
	sig.Hash = data[:64]
	length := binary.LittleEndian.Uint64(data[64:72])
	end := 72 + int(length)
	sig.Sig = data[72:end]
	sig.Type = string(data[end:])
	return sig, nil
}

func EncodeSignature(out io.Writer, sig *Signature) error {
	data, err := serialise(sig)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "STARSIGN SIGNATURE",
		Bytes: data,
	}
	if err = pem.Encode(out, block); err != nil {
		return err
	}
	return nil
}

func DecodeSignature(in io.Reader) (sig *Signature, err error) {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			sig = nil
			err = fmt.Errorf("Malformed signature data")
		}
	}()
	block, _ := pem.Decode(data)
	if block.Type != "STARSIGN SIGNATURE" {
		return nil, fmt.Errorf("Not a Starsign signature: %s", block.Type)
	}
	if len(block.Headers) != 0 {
		return nil, fmt.Errorf("Signature contains untrusted headers")
	}
	return deserialise(block.Bytes)
}

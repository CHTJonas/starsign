package starsign

import (
	"bytes"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
)

const pemType = "STARSIGN SIGNATURE"
const blakeHashLength = 64
const sigLenLength = 8

var ErrMalformedSig = errors.New("Malformed signature data")
var ErrSigDataTooShort = errors.New("Signature data too short")
var ErrNotStarsignSig = errors.New("Not a Starsign signature")
var ErrUntrustHeaders = errors.New("Signature contains untrusted headers")

func serialise(sig *Signature) ([]byte, error) {
	if len(sig.Hash) != blakeHashLength {
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
	sigLenEnd := blakeHashLength + sigLenLength
	if len(data) <= sigLenEnd {
		return nil, ErrSigDataTooShort
	}
	sig := new(Signature)
	sig.Hash = data[:blakeHashLength]
	length := binary.LittleEndian.Uint64(data[blakeHashLength:sigLenEnd])
	sigEnd := sigLenEnd + int(length)
	if len(data) <= sigEnd {
		return nil, ErrSigDataTooShort
	}
	sig.Sig = data[sigLenEnd:sigEnd]
	sig.Type = string(data[sigEnd:])
	return sig, nil
}

func EncodeSignature(out io.Writer, sig *Signature) error {
	data, err := serialise(sig)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  pemType,
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
			err = ErrMalformedSig
		}
	}()
	block, _ := pem.Decode(data)
	if block.Type != pemType {
		return nil, ErrNotStarsignSig
	}
	if len(block.Headers) != 0 {
		return nil, ErrUntrustHeaders
	}
	return deserialise(block.Bytes)
}

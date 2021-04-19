package starsign

import (
	"encoding/pem"
	"errors"
	"io"
)

const pemType = "STARSIGN SIGNATURE"
const blakeHashLength = 64
const sigLenLength = 8

var ErrHashWrongLength = errors.New("unexpected BLAKE2 hash length")
var ErrSigDataTooShort = errors.New("signature data too short")
var ErrMalformedSig = errors.New("malformed signature data")
var ErrNotStarsignSig = errors.New("not a Starsign signature")
var ErrUntrustHeaders = errors.New("signature contains untrusted headers")

func serialise(sig *Signature) ([]byte, error) {
	if len(sig.Hash) != blakeHashLength {
		return nil, ErrHashWrongLength
	}
	return append(sig.Hash, sig.Sig...), nil
}

func deserialise(data []byte) (*Signature, error) {
	if len(data) <= blakeHashLength {
		return nil, ErrSigDataTooShort
	}
	sig := new(Signature)
	sig.Hash = data[:blakeHashLength]
	sig.Sig = data[blakeHashLength:]
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
	data, err := io.ReadAll(in)
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

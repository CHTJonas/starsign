package starsign

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type Signature struct {
	Hash []byte
	Sig  []byte
	Type string
}

func Sign(in io.Reader) (*Signature, error) {
	socket := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, fmt.Errorf("Failed to open SSH_AUTH_SOCK: %v", err)
	}
	client := agent.NewClient(conn)
	keys, err := client.List()
	if err != nil {
		return nil, fmt.Errorf("Failed to list SSH keys: %v", err)
	}
	hash, err := hash(in)
	if err != nil {
		return nil, fmt.Errorf("Failed to hash data: %v", err)
	}
	sig, err := client.Sign(keys[0], hash)
	if err != nil {
		return nil, fmt.Errorf("Failed to sign data: %v", err)
	}
	return &Signature{
		Hash: hash,
		Sig:  sig.Blob,
		Type: sig.Format,
	}, nil
}

func Verify(in io.Reader, sig *Signature, key ssh.PublicKey) error {
	hash, err := hash(in)
	if err != nil {
		return fmt.Errorf("Failed to hash data: %v", err)
	}
	if !bytes.Equal(hash, sig.Hash) {
		return fmt.Errorf("Hash mismatch")
	}
	sshSig := &ssh.Signature{
		Format: sig.Type,
		Blob:   sig.Sig,
	}
	return key.Verify(hash[:], sshSig)
}

func hash(in io.Reader) ([]byte, error) {
	hasher, err := blake2b.New512(nil)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(hasher, in)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

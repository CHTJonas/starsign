package starsign

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

var ErrPubKeyNotFound = errors.New("public key not found in agent")
var ErrHashMismatch = errors.New("hash mismatch")

type Signature struct {
	Hash []byte
	Sig  []byte
}

func Sign(in io.Reader, pubKey ssh.PublicKey) (*Signature, error) {
	socket := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, fmt.Errorf("failed to open SSH_AUTH_SOCK: %w", err)
	}
	client := agent.NewClient(conn)
	keys, err := client.List()
	if err != nil {
		return nil, fmt.Errorf("failed to list SSH keys: %w", err)
	}
	key := keys[0]
	if pubKey != nil {
		key = nil
		pubKeyBytes := pubKey.Marshal()
		for _, k := range keys {
			if bytes.Equal(k.Marshal(), pubKeyBytes) {
				key = k
				break
			}
		}
		if key == nil {
			return nil, ErrPubKeyNotFound
		}
	}
	hash, err := hash(in)
	if err != nil {
		return nil, fmt.Errorf("failed to hash data: %w", err)
	}
	sig, err := client.Sign(key, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign data: %w", err)
	}
	return &Signature{
		Hash: hash,
		Sig:  sig.Blob,
	}, nil
}

func Verify(in io.Reader, sig *Signature, key ssh.PublicKey) error {
	sshSig := &ssh.Signature{
		Format: key.Type(),
		Blob:   sig.Sig,
	}
	if err := key.Verify(sig.Hash, sshSig); err != nil {
		return fmt.Errorf("failed to verify signature: %w", err)
	}
	hash, err := hash(in)
	if err != nil {
		return fmt.Errorf("failed to hash data: %w", err)
	}
	if !bytes.Equal(hash, sig.Hash) {
		return ErrHashMismatch
	}
	return nil
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

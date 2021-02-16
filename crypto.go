package starsign

import (
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type Signature struct {
	Hash [blake2b.Size]byte
	Sig  []byte
	Type string
}

func Sign(data []byte) *Signature {
	socket := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", socket)
	if err != nil {
		log.Fatalf("Failed to open SSH_AUTH_SOCK: %v", err)
	}
	client := agent.NewClient(conn)
	keys, err := client.List()
	if err != nil {
		log.Fatalf("Failed to list SSH keys: %v", err)
	}
	hash := blake2b.Sum512(data)
	sig, err := client.Sign(keys[0], hash[:])
	if err != nil {
		log.Fatalf("Failed to sign: %v", err)
	}
	return &Signature{
		Hash: hash,
		Sig:  sig.Blob,
		Type: sig.Format,
	}
}

func Verify(data []byte, sig *Signature, key ssh.PublicKey) error {
	hash := blake2b.Sum512(data)
	if hash != sig.Hash {
		return fmt.Errorf("Hash mismatch")
	}
	sshSig := &ssh.Signature{
		Format: sig.Type,
		Blob:   sig.Sig,
	}
	return key.Verify(hash[:], sshSig)
}

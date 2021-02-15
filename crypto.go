package starsign

import (
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func Sign(data []byte) *ssh.Signature {
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

	sig, err := client.Sign(keys[0], data)
	if err != nil {
		log.Fatalf("Failed to sign data: %v", err)
	}
	return sig
}

func Verify(data []byte, sig *ssh.Signature, key ssh.PublicKey) error {
	return key.Verify(data, sig)
}

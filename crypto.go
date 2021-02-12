package starsign

import (
	"encoding/base64"
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

func Verify(data []byte, sig *ssh.Signature) error {
	key := dummyTestingKey()
	return key.Verify(data, sig)
}

func dummyTestingKey() *agent.Key {
	fingerprint := "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCsu/KmxxHvrQy4OorfEqF5zLfxk/QFDYs2MweLCvZjhkvUr6xKV6GXYH3W5Rq6BSKIzj3qqAB9yZ5G5oXXEjPs="
	blob, err := base64.StdEncoding.DecodeString(fingerprint)
	if err != nil {
		log.Fatalf("Failed to deserialise key: %v", err)
	}
	return &agent.Key{
		Format:  "ecdsa-sha2-nistp256",
		Blob:    blob,
		Comment: "yubikey-5-nfc",
	}
}

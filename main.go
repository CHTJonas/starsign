package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func main() {
	data := []byte("testing")
	sig := sign(data)
	out := fmt.Sprintf("%s %s", sig.Format, base64.StdEncoding.EncodeToString(sig.Blob))
	fmt.Println(out)
	if err := verify(data, sig); err != nil {
		fmt.Println("!!! SIGNATURE VERIFICATION FAILED !!!")
	} else {
		fmt.Println("Signature ok")
	}
}

func sign(data []byte) *ssh.Signature {
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

func verify(data []byte, sig *ssh.Signature) error {
	fingerprint := "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCsu/KmxxHvrQy4OorfEqF5zLfxk/QFDYs2MweLCvZjhkvUr6xKV6GXYH3W5Rq6BSKIzj3qqAB9yZ5G5oXXEjPs="
	blob, err := base64.StdEncoding.DecodeString(fingerprint)
	if err != nil {
		log.Fatalf("Failed to deserialise key: %v", err)
	}
	key := &agent.Key{
		Format:  "ecdsa-sha2-nistp256",
		Blob:    blob,
		Comment: "yubikey-5-nfc",
	}
	return key.Verify(data, sig)
}

package starsign

import (
	"bufio"
	"io"
	"os"

	"golang.org/x/crypto/ssh"
)

func ReadPubKeyFile(in io.Reader) (ssh.PublicKey, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	key, _, _, _, err := ssh.ParseAuthorizedKey(data)
	return key, err
}

func GetFileWriter(path string) (*bufio.Writer, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return bufio.NewWriter(file), nil
}

func GetFileReader(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(file), nil
}

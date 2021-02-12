package starsign

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func ReadPubKeyFile(path string) (ssh.PublicKey, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	key, _, _, _, err := ssh.ParseAuthorizedKey(data)
	return key, err
}

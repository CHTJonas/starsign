package main

import (
	"encoding/base64"
	"fmt"

	"github.com/CHTJonas/starsign"
)

func main() {
	data := []byte("testing")
	sig := starsign.Sign(data)
	out := fmt.Sprintf("%s %s", sig.Format, base64.StdEncoding.EncodeToString(sig.Blob))
	fmt.Println(out)

	s, _ := starsign.EncodeSignature(sig)
	sig, _ = starsign.DecodeSignature(s)
	key, _ := starsign.ReadPubKeyFile("/Users/charlie/.ssh/yk.pub")

	if err := starsign.Verify(data, key, sig); err != nil {
		fmt.Println("!!! SIGNATURE VERIFICATION FAILED !!!")
	} else {
		fmt.Println("Signature ok")
	}
}
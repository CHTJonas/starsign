package main

import (
	"fmt"

	"github.com/CHTJonas/starsign"
)

func main() {
	data := []byte("testing")

	sig := starsign.Sign(data)
	bufW, _ := starsign.GetFileWriter("sig.txt")
	starsign.EncodeSignature(bufW, sig)
	bufW.Flush()

	key, _ := starsign.ReadPubKeyFile("/Users/charlie/.ssh/yk.pub")
	bufR, _ := starsign.GetFileReader("sig.txt")
	sig, _ = starsign.DecodeSignature(bufR)

	if err := starsign.Verify(data, key, sig); err != nil {
		fmt.Println("!!! SIGNATURE VERIFICATION FAILED !!!")
	} else {
		fmt.Println("Signature ok")
	}
}

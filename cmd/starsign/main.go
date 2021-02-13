package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CHTJonas/starsign"
)

var (
	// Software version defaults to the value below but is overridden by the compiler in Makefile.
	version     = "dev-edge"
	versionFlag bool
	licenseFlag bool
	signFlag    bool
	verifyFlag  bool
	armorFlag   bool
	outFlag     string
)

func init() {
	flag.BoolVar(&versionFlag, "V", false, "print the version")
	flag.BoolVar(&versionFlag, "version", false, "print the version")
	flag.BoolVar(&signFlag, "s", false, "generate a digital signature")
	flag.BoolVar(&signFlag, "sign", false, "generate a digital signature")
	flag.BoolVar(&verifyFlag, "v", false, "verify a digital signature")
	flag.BoolVar(&verifyFlag, "verify", false, "verify a digital signature")
	flag.BoolVar(&armorFlag, "a", false, "generate armored output")
	flag.BoolVar(&armorFlag, "armor", false, "generate armored output")
	flag.StringVar(&outFlag, "o", "", "output to `FILE` (default is input filename with .star extension)")
	flag.StringVar(&outFlag, "output", "", "output to `FILE` (default is input filename with .star extension)")
	flag.Parse()
}

func main() {
	switch {
	case versionFlag:
		fmt.Println(version)
		os.Exit(0)
	case licenseFlag:
		fmt.Println("Copyright (c) 2021 Charlie Jonas.")
		fmt.Println("This software is released under the BSD 2-Clause License.")
		fmt.Println("Please visit https://github.com/CHTJonas/starsign for more information.")
		os.Exit(0)
	case signFlag:
		if err := sign(); err != nil {
			fmt.Println("Failed to generate signature")
			os.Exit(125)
		}
		fmt.Println("Signature generated")
		os.Exit(0)
	case verifyFlag:
		if err := verify(); err != nil {
			fmt.Println("!!! SIGNATURE VERIFICATION FAILED !!!")
			os.Exit(125)
		}
		fmt.Println("Signature ok")
		os.Exit(0)
	default:
		fmt.Println("Must specify one of either -s/--sign or -v/--verify")
		os.Exit(1)
	}
}

var data = []byte("testing")

func sign() error {
	sig := starsign.Sign(data)
	bufW, err := starsign.GetFileWriter("sig.txt")
	if err != nil {
		return err
	}
	defer bufW.Flush()
	return starsign.EncodeSignature(bufW, sig)
}

func verify() error {
	key, err := starsign.ReadPubKeyFile("/Users/charlie/.ssh/yk.pub")
	if err != nil {
		return err
	}
	bufR, err := starsign.GetFileReader("sig.txt")
	if err != nil {
		return err
	}
	sig, err := starsign.DecodeSignature(bufR)
	if err != nil {
		return err
	}
	return starsign.Verify(data, key, sig)
}

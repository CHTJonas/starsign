package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/CHTJonas/starsign"
)

var (
	// Software version defaults to the value below but is overridden by the compiler in Makefile.
	version = "dev-edge"

	// Command line flags.
	versionFlag bool
	licenseFlag bool
	signFlag    bool
	verifyFlag  bool
	pubKeyFlag  string
	outFlag     string
)

func init() {
	flag.Usage = func() {
		fmt.Println(usage)
	}
	flag.BoolVar(&versionFlag, "V", false, "print the version")
	flag.BoolVar(&versionFlag, "version", false, "print the version")
	flag.BoolVar(&signFlag, "s", false, "generate a digital signature")
	flag.BoolVar(&signFlag, "sign", false, "generate a digital signature")
	flag.BoolVar(&verifyFlag, "v", false, "verify a digital signature")
	flag.BoolVar(&verifyFlag, "verify", false, "verify a digital signature")
	flag.StringVar(&pubKeyFlag, "p", "", "public key file to use when verifying")
	flag.StringVar(&pubKeyFlag, "pubkey", "", "public key file to use when verifying")
	flag.StringVar(&outFlag, "o", "", "file to write to when signing")
	flag.StringVar(&outFlag, "output", "", "file to write to when signing")
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

func openFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open file", path+":", err)
		os.Exit(99)
	}
	return f
}

func createFile(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Failed to create file", path+":", err)
		os.Exit(99)
	}
	return f
}

func sign() error {
	var (
		in  io.Reader
		out io.Writer
	)

	inPath := flag.Arg(0)
	if flag.NArg() != 1 || inPath == "" {
		fmt.Println("Error: wrong number of arguments.")
		fmt.Println("In signature mode Starsign accepts a single argument to specify the input file.")
		os.Exit(1)
	}
	if inPath == "-" {
		if outFlag == "" {
			outFlag = "-"
		}
		in = os.Stdin
	} else {
		if outFlag == "" {
			outFlag = inPath + ".starsig"
		}
		f := openFile(inPath)
		defer f.Close()
		in = f
	}
	if outFlag == "-" {
		out = os.Stdout
	} else {
		f := createFile(outFlag)
		defer f.Close()
		out = f
	}

	sig, err := starsign.Sign(in)
	if err != nil {
		return err
	}
	return starsign.EncodeSignature(out, sig)
}

func verify() error {
	if flag.NArg() < 1 || flag.NArg() > 2 {
		fmt.Println("Error: wrong number of arguments.")
		fmt.Println("In verification mode Starsign accepts a single argument to specify the input file and an optional argument to the signature file.")
		os.Exit(1)
	}
	if pubKeyFlag == "" {
		fmt.Println("Error: public key file not specified.")
		os.Exit(1)
	}

	var d, s, k *os.File
	dataPath := flag.Arg(0)
	d = openFile(dataPath)
	defer d.Close()
	if sigPath := flag.Arg(1); sigPath != "" {
		s = openFile(sigPath)
		defer s.Close()
	} else {
		s = openFile(dataPath + ".starsig")
		defer s.Close()
	}
	k = openFile(pubKeyFlag)
	defer k.Close()

	key, err := starsign.ReadPubKeyFile(k)
	if err != nil {
		return err
	}
	sig, err := starsign.DecodeSignature(s)
	if err != nil {
		return err
	}
	return starsign.Verify(d, sig, key)
}

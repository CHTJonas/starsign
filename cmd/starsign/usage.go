package main

const usage = `Usage:
    starsign -s [-p PUBLICKEY] [-o OUTPUT] INPUT
    starsign -v -p PUBLICKEY INPUT [SIGFILE]

Mode flags:
    -s, --sign                  Sign a file
    -v, --verify                Verify a signature
    -V, --version               Print version and exit
    -l, --license               Print license and exit

Option flags:
    -o, --output OUTPUT         Write the signature to the file at path OUTPUT.
    -p, --pubkey PUBLICKEY      Use the public key from the file at path PUBLICKEY.
    --verbose                   Print verbose error messages when signing or verifying.

In sign mode, Starsign takes an argument INPUT that determines the file to sign.
This can be "-" in which case the input will be taken from the standard input. In
the event that the "-o OUTPUT" flag is not used then OUTPUT will default to INPUT
but with the suffix ".starsig" appended.

In verify mode, Starsign takes an INPUT argument which determins the file that has
been signed. Starsign assumes that SIGFILE is the same as INPUT but with the
".starsig" suffix appended, unless SIGFILE has been specified manually.

Example usage:
    $ starsign -s source.tar.gz
    $ echo 'testing' | starsign -s -
    $ starsign -s -o message.sig message.txt
    $ starsign -v -p ~/.ssh/id_rsa.pub source.tar.gz
    $ starsign -v -p ~/.ssh/id_rsa.pub message.txt message.sig`

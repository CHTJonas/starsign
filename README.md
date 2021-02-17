# Starsign

Starsign is the world's simplest tool to sign and verify file signatures to ensure authenticity.

There's already a myriad of other crypto tools and libraries out there for managing digital signatures so why choose Starsign? Well, here are my motivations:

* I **don't** want to use anything PGP-based.
* I want a tool that's written in a modern memory-safe language.
* I want a tool that is easy to audit and has minimal third-party dependencies (i.e. StdLib only).
* I want a tool that's totally idiot-proof. There should be one option to sign and one to verify. Nothing else.
* I [already distribute](https://github.com/CHTJonas.keys) my SSH keys so let's reuse those for signing.
* I use a [YubiKey](https://www.yubico.com/) with Filippo Valsorda's [SSH Agent](https://filippo.io/yubikey-agent) so let's integrate nicely with that.

If you think those all sound like good ideas then Starsign may be for you.

## Usage

```
Usage:
    starsign -s [-o OUTPUT] INPUT
    starsign -v -p PUBLICKEY INPUT [SIGFILE]

Mode flags:
    -s, --sign                  Sign a file
    -v, --verify                Verify a signature
    -V, --version               Print version and exit
    -l, --license               Print license and exit

Option flags:
    -o, --output OUTPUT         Write the signature to the file at path OUTPUT.
    -p, --pubkey PUBLICKEY      Read in public key from the file at path PUBLICKEY.
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
    $ starsign -v -p ~/.ssh/id_rsa.pub message.txt message.sig
```

## Compiling

To build Starsign from source you will need a suitable [Go toolchain installed](https://golang.org/doc/install). After that just clone the project using Git and run Make! Cross-compilation is easy in Go so by default we compile binaries for various operating systems and processor architectures and place them all in `./bin`:

```bash
git checkout https://github.com/CHTJonas/starsign.git
cd starsign
make clean && make all
```

---

### Copyright

Starsign is licensed under the [BSD 2-Clause License](https://opensource.org/licenses/BSD-2-Clause).

Copyright (c) 2021 Charlie Jonas.

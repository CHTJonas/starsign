# Starsign

Starsign is the world's simplest tool to sign and verify file signatures to ensure authenticity. It generates a 64-byte BLAKE2 hash of the input data and signs this via your SSH agent, using either the first SSH key that's loaded or a specific key who's public key path you pass on the command line.

## Motivations

There's already a myriad of other cryptographic tools and libraries out there for managing digital signatures so why use something new? Here's a few of my goals for Starsign:

* I want a tool that's as simple to use as possible with one command line flag to sign and one to verify.
* I want a tool that's composable and works well with UNIX pipes or as a library for other Go applications.
* I want a tool that has little, if any, configuration.
* I want a tool that reuses my [already distributed SSH keys](https://github.com/CHTJonas.keys) for signing.
* I want a tool that is easy to audit and has minimal third-party dependencies (i.e. standard library only).
* I want a tool that's written in a modern memory-safe language and which is easy to cross-compile.
* I **don't** want to use anything PGP-based.

If those all seem like good ideas then Starsign may be for you.

## Usage

Make sure that the `SSH_AUTH_SOCK` environment variable contains the path to your agent's UNIX socket. Personally I use Filippo Valsorda's [yubikey-agent](https://github.com/FiloSottile/yubikey-agent) but it also works perfectly with OpenSSH's agent.

```
Usage:
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
    $ starsign -v -p ~/.ssh/id_rsa.pub message.txt message.sig
```

## Installation

Pre-built binaries for a variety of operating systems and architectures are available to download from [GitHub Releases](https://github.com/CHTJonas/starsign/releases). I will generate a `SHA256SUMS` file and sign it using the SSH key that resides on my YubiKey and which has the following fingerprint (which you can cross-check [here](https://chtj2.user.srcf.net/identity/authorized_keys) and [here](https://github.com/CHTJonas.keys)):

```
ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCsu/KmxxHvrQy4OorfEqF5zLfxk/QFDYs2MweLCvZjhkvUr6xKV6GXYH3W5Rq6BSKIzj3qqAB9yZ5G5oXXEjPs=
```

To compile Starsign from source you will need a suitable [Go toolchain installed](https://golang.org/doc/install). After that just clone the project using Git and run Make! Cross-compilation is easy in Go so by default we build for all targets and place the resulting executables in `./bin`:

```bash
git checkout https://github.com/CHTJonas/starsign.git
cd starsign
make clean && make all
```

## Copyright

Starsign is licensed under the [BSD 2-Clause License](https://opensource.org/licenses/BSD-2-Clause).

Copyright (c) 2021 Charlie Jonas.

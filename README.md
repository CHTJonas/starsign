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

`TODO`

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

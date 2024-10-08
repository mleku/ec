github.com/mleku/ec
=====

zap me to support my work ⚡️mleku@getalby.com

This is a full drop-in replacement for
[github.com/btcsuite/btcd/btcec](https://github.com/btcsuite/btcd/tree/master/btcec)
eliminating the import from the Decred repository, and including the chainhash
helper functions, needed for hashing messages for signatures.

This includes the old style compact secp256k1 ECDSA signatures, that recover the 
public key rather than take a key as a parameter as used in Bitcoin
transactions, the new style Schnorr signatures, and the Musig2 implementation.

BIP 340 Schnorr signatures are implemented except for the nonstandard message length
tests, that nobody uses anyway.

The remainder of this document is from the original README.md.

------------------------------------------------------------------------------

Package `ec` implements elliptic curve cryptography needed for working with
Bitcoin. It is designed so that it may be used with the standard
crypto/ecdsa packages provided with Go.

A comprehensive suite of testis provided to ensure proper functionality.

Package btcec was originally based on work from ThePiachu which is licensed under
the same terms as Go, but it has signficantly diverged since then. The btcsuite
developers original is licensed under the liberal ISC license.

## Installation and Updating

```bash
$ go install -u -v github.com/mleku/ec
```

## dusk-crypto: Cryptography Module
### Overview
#### BN-256
Optimal ate pairing over 256 bit Barretto-Naehrig curves providing approximately 96 bits of security [1]. Dusk Network BN-256 library can be found [here](https://github.com/dusk-network/bn256).

#### BLS
Boneh-Lynn-Shacham signatures use BN256 to provide a short aggregateable signature scheme [2]. Trivially since signatures are elements of an elliptic curve, they are short and since elliptic curves do not have a ring structure, they are not susceptible to index-calculus attacks. Dusk Network BLS library provides the full implementation of the next generation BLS Multisignature with Public Key aggregation [3]. This improves on the the previous algorithm by providing:

* protection versus rogue-key attack. 
* aggregation of multiple public keys. 
* aggregation of multiple signatures. 
* a method for hashing to the curve.
* (multi-) signature compression and compression verification

#### bLSAG
A linkable ring signature scheme whose security is based on the Discrete Logarithm Problem [4]. The signature size grows linearly with the number of members in the ring. This is a zero knowledge proof where we prove that at most one member from the ring has signed a given message from the provided public keys, without revealing which member has signed.

#### Range Proof
A proof that an element x is within a discrete set [0, 2^N], where in our case N is 64. This is a zero knowledge proof, where we prove that this element is within the given range without providing any extra information. This specific rangeproof uses the Bulletproof protocol [5], which uses a inner profuct proof of knowledge to compress the final vectors. Due to the inner product, the rangeproof grows logarithmically with N.

### References
[1] Naehrig, M.; Niederhagen, R.; Schwabe, P. (2010). New software speed records for cryptographic pairings. Link:
https://cryptojedi.org/papers/dclxvi-20100714.pdf

[2] Boneh, D.; Lynn, B.; Shacham, H. (2016). Short signatures from the Weil pairing. Link: https://www.iacr.org/archive/asiacrypt2001/22480516.pdf

[3] Boneh, D.; Drijvers, M.; Neven, G. (2018). Compact Multi-Signatures for Smaller Blockchains. Link: https://eprint.iacr.org/2018/483.pdf

[4] Back, A. (2015). Ring signature efficiency. Link: https://bitcointalk.org/index.php?topic=972541

[5] Bunz, B.; Bootle, J.; Boneh, D.; Poelstra, A.; Wuille, P.; Maxwell, G. (2017). Bulletproofs: Short Proofs for Confidential Transactions and More. Link: https://eprint.iacr.org/2017/1066.pdf

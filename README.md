## dusk-crypto: Cryptography Module
### Overview
#### BN-256
Optimal ate pairing over 256 bit Barretto-Naehrig curves providing approximately 96 bits of security [1]. Dusk Network BN-256 library has been forked from the Golang Crypto [repository]( https://github.com/golang/crypto/tree/master/bn256) with an addition of point compression method.

#### BLS
Boneh-Lynn-Shacham signatures use BN256 to provide a short aggregateable signature scheme [2]. Trivially since signatures are elements of an elliptic curve, they are short and since elliptic curves do not have a ring structure, they are not susceptible to index-calculus attacks. Dusk Network BLS library has been forked from the Vuvuzela Crypto [repository](https://github.com/vuvuzela/crypto/blob/master/bls) with additions of rogue-key attack mitigation, public key aggregation and a hash to curve methods.

#### bLSAG
A linkable ring signature scheme whose security is based on the Discrete Logarithm Problem [3]. The signature size grows linearly with the number of members in the ring. This is a zero knowledge proof where we prove that at most one member from the ring has signed a given message from the provided public keys, without revealing which member has signed.

#### Range Proof
A proof that an element x is within a discrete set [0, 2^N], where in our case N is 64. This is a zero knowledge proof, where we prove that this element is within the given range without providing any extra information. This specific rangeproof uses the Bulletproof protocol [4], which uses a inner profuct proof of knowledge to compress the final vectors. Due to the inner product, the rangeproof grows logarithmically with N.

### References
[1] Naehrig, M.; Niederhagen, R.; Schwabe, P. (2010). New software speed records for cryptographic pairings. Link:
https://cryptojedi.org/papers/dclxvi-20100714.pdf

[2] Boneh, D.; Lynn, B.; Shacham, H. (2016). Short signatures from the Weil pairing. Link: https://www.iacr.org/archive/asiacrypt2001/22480516.pdf

[3] Back, A. (2015). Ring signature efficiency. Link: https://bitcointalk.org/index.php?topic=972541

[4] Bunz, B.; Bootle, J.; Boneh, D.; Poelstra, A.; Wuille, P.; Maxwell, G. (2017). Bulletproofs: Short Proofs for Confidential Transactions and More. Link: https://eprint.iacr.org/2017/1066.pdf

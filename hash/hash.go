package hash

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"hash"
	"strconv"

	"github.com/OneOfOne/xxhash"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

// Hash is a convenient alias of hash.Hash
type Hash = hash.Hash

// Blake2b256 performs a Blake2 hashing of a binary payload
func Blake2b256(bs []byte) ([]byte, error) {
	h, err := blake2b.New256(nil)
	if err != nil {
		return nil, err
	}
	return PerformHash(h, bs)
}

// Blake2b512 performs a Blake2 hashing of a binary payload
func Blake2b512(bs []byte) ([]byte, error) {
	h, err := blake2b.New512(nil)
	if err != nil {
		return nil, err
	}
	return PerformHash(h, bs)
}

// Sha3256 takes a byte slice
// and returns the SHA3-256 hash
func Sha3256(bs []byte) ([]byte, error) {
	return PerformHash(sha3.New256(), bs)
}

// PerformHash takes a generic hash.Hash and returns the hashed payload
func PerformHash(H Hash, bs []byte) ([]byte, error) {
	_, err := H.Write(bs)
	if err != nil {
		return nil, err
	}
	return H.Sum(nil), err
}

// Sha3512 takes a byte slice
// and returns the SHA3-512 hash
func Sha3512(bs []byte) ([]byte, error) {
	return PerformHash(sha3.New512(), bs)
}

// Xxhash uses 64 bit xxhashing. It is not meant as a trapdoor, but as a fast collision resistant hashing for performant equal check and store retrieval
func Xxhash(bs []byte) ([]byte, error) {
	return PerformHash(xxhash.New64(), bs)
}

// RandEntropy takes a uint32 argument n and populates a byte slice of
// size n with random input.
func RandEntropy(n uint32) ([]byte, error) {

	b := make([]byte, n)
	a, err := rand.Read(b)

	if err != nil {
		return nil, errors.New("Error generating entropy " + err.Error())
	}
	if uint32(a) != n {
		return nil, errors.New("Error expected to read" + strconv.Itoa(int(n)) + " bytes instead read " + strconv.Itoa(a) + " bytes")
	}
	return b, nil
}

//CompareChecksum takes data and an expected checksum
// Returns true if the checksum of the given data is
// equal to the expected checksum
func CompareChecksum(data []byte, want uint32) bool {
	got, err := Checksum(data)
	if err != nil {
		return false
	}
	if got != want {
		return false
	}
	return true
}

// Checksum hashes the data with Xxhash
// and returns the first four bytes
func Checksum(data []byte) (uint32, error) {
	hash, err := Xxhash(data)
	if err != nil {
		return 0, err
	}
	checksum := binary.BigEndian.Uint32(hash[:4])
	return checksum, err
}

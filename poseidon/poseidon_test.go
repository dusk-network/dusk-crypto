package poseidon

import (
	"testing"

	ristretto "github.com/bwesterb/go-ristretto"
	"github.com/stretchr/testify/assert"
)

func TestPoseidonHashOne(t *testing.T) {
	p := New()

	p.WriteScalar(ristretto.Scalar{1, 0, 0, 0, 0, 0, 0, 0})
	p.WriteScalar(ristretto.Scalar{1, 0, 0, 0, 0, 0, 0, 0})
	p.WriteScalar(ristretto.Scalar{1, 0, 0, 0, 0, 0, 0, 0})
	p.WriteScalar(ristretto.Scalar{1, 0, 0, 0, 0, 0, 0, 0})

	digest := []byte{0x0d, 0x03, 0x47, 0xfc, 0xa6, 0x5b, 0xea, 0xac, 0x6f, 0x5f, 0x8d, 0xea, 0xc3, 0x42, 0xe7, 0xd6, 0xfd, 0x85, 0x9f, 0xfa, 0x1c, 0x79, 0x89, 0x52, 0xf8, 0xdb, 0x1b, 0x19, 0xd5, 0x25, 0xb5, 0x0}
	assert.Equal(t, digest, p.Sum(nil))
}

func TestQuinticSbox(t *testing.T) {
	a := ristretto.Scalar{0xe2d76bf9, 0xbb6e333c, 0x2ec4e479, 0xba272f09, 0x046d4aca, 0x6aadbd72, 0x95c9842a, 0x0f0cdba9}
	b := ristretto.Scalar{0xe41683eb, 0xc0f550ab, 0x7f547e18, 0x935175b2, 0xf72488bf, 0x03384905, 0x30658415, 0x0cf11c8e}

	QuinticSbox(&a)
	assert.Equal(t, b, a)
}

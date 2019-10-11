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

	digest := []byte{0x15, 0xf4, 0x2e, 0x34, 0xae, 0x8f, 0x31, 0x2e, 0xb8, 0xa3, 0x3b, 0x28, 0x30, 0xd1, 0x94, 0x8d, 0xcf, 0x06, 0xd3, 0xa2, 0xeb, 0x65, 0xef, 0xf8, 0x65, 0x2e, 0x78, 0x08, 0xcf, 0xc0, 0xff, 0x03}
	assert.Equal(t, digest, p.Sum(nil))
}

func TestQuinticSbox(t *testing.T) {
	a := ristretto.Scalar{0xe2d76bf9, 0xbb6e333c, 0x2ec4e479, 0xba272f09, 0x046d4aca, 0x6aadbd72, 0x95c9842a, 0x0f0cdba9}
	b := ristretto.Scalar{0xe41683eb, 0xc0f550ab, 0x7f547e18, 0x935175b2, 0xf72488bf, 0x03384905, 0x30658415, 0x0cf11c8e}

	QuinticSbox(&a)
	assert.Equal(t, b, a)
}

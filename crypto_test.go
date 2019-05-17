package crypto_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"dusk-crypto/dusk-go/pkg/crypto"
)

func TestRandEntropy(t *testing.T) {

	for i := 0; i < 100; i++ {
		n := uint32(rand.Intn(1e3))
		en, err := crypto.RandEntropy(n)

		assert.Equal(t, nil, err)
		assert.Equal(t, uint32(len(en)), uint32(n))

	}
}

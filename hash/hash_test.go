package hash

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randomMessage(size int) []byte {
	msg := make([]byte, size)
	_, _ = rand.Read(msg)
	return msg
}

func BenchmarkXxhash(b *testing.B) {

	testBytes := randomMessage(32)

	for i := 0; i < b.N; i++ {
		_, _ = Xxhash(testBytes)
	}
}

func BenchmarkSha3(b *testing.B) {

	testBytes := randomMessage(32)

	for i := 0; i < b.N; i++ {
		_, _ = Sha3256(testBytes)
	}
}

func TestRandEntropy(t *testing.T) {

	for i := 0; i < 100; i++ {
		n := uint32(rand.Intn(1e3))
		en, err := RandEntropy(n)

		assert.Equal(t, nil, err)
		assert.Equal(t, uint32(len(en)), uint32(n))

	}
}

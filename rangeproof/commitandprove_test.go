package rangeproof

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlow(t *testing.T) {

	n := rand.Int63()
	c, errc := Commit(n)
	p, errp := GenProof(n, c)
	ok, errv := VerifyProof(p)
	require.Nil(t, errc)
	require.Nil(t, errp)
	require.Nil(t, errv)
	assert.Equal(t, true, ok)
}

package rangeproof

import (
	"math/big"

	ristretto "github.com/bwesterb/go-ristretto"
	"github.com/vosbor/dusk-crypto/rangeproof/pedersen"
)

type Commitment struct {
	PedersenCommitment pedersen.Commitment
}

type RangeProof struct {
	P Proof
}

func Commit(v int64) (Commitment, error) {

	genData := []byte("vosbor.BulletProof.v1")
	ped := pedersen.New(genData)
	ped.BaseVector.Compute(uint32((M * N)))

	var amount ristretto.Scalar
	amount.SetBigInt(big.NewInt(v))
	c := ped.CommitToScalar(amount)

	output := Commitment{
		PedersenCommitment: c,
	}

	return output, nil
}

func GenProof(v int64, c Commitment) (RangeProof, error) {
	amounts := []ristretto.Scalar{}
	commitments := make([]pedersen.Commitment, 0, M)

	// N is number of bits in range
	// So amount will be between 0...2^(N-1)
	const N = 64

	var amount ristretto.Scalar
	amount.SetBigInt(big.NewInt(v))
	amounts = append(amounts, amount)
	commitments = append(commitments, c.PedersenCommitment)

	p, err := Prove(amounts, commitments, true)

	output := RangeProof{
		P: p,
	}

	return output, err

}

func VerifyProof(p RangeProof) (bool, error) {
	ok, err := Verify(p.P)
	return ok, err
}

package poseidon

import (
	"errors"

	ristretto "github.com/bwesterb/go-ristretto"
)

// Poseidon represents the structure for the main hash unit
type Poseidon struct {
	params *Params
	input  []ristretto.Scalar
}

// New will generate an instance of Poseidon with Default parameters
func New() Poseidon {
	params := Default()
	return NewWithParams(&params)
}

// NewWithParams returns an instance of Poseidon with the specified parameters
func NewWithParams(params *Params) Poseidon {
	width := uint32(1) << uint32(params.Width-1)
	firstElementScalar := ristretto.Scalar{width - 1}

	return Poseidon{params: params, input: []ristretto.Scalar{firstElementScalar}}
}

// Size will return the maximum allowed length for input elements
func (p *Poseidon) Size() int {
	return p.params.Width
}

// BlockSize ..
func (p *Poseidon) BlockSize() int {
	// TODO not implemented
	return 0
}

// Write will convert the bytes to Scalar with a reduction, and append the result to the inner structure
func (p *Poseidon) Write(b []byte) (int, error) {
	s := ristretto.Scalar{}
	s.Derive(b)

	return p.WriteScalar(s)
}

// WriteScalar will try to append the provided scalar to the input set
func (p *Poseidon) WriteScalar(s ristretto.Scalar) (int, error) {
	if p.params.Width == len(p.input) {
		return 0, errors.New("Maximum width reached")
	}

	p.input = append(p.input, s)

	return len(s), nil
}

// Sum will compute the Poseidon digest value. The usage of the bytes parameter is currently not implemented
func (p *Poseidon) Sum(in []byte) []byte {
	p.Pad()

	keysOffset := 0

	for i := 0; i < p.params.FullRoundsBeginning; i++ {
		p.applyFullRound(&keysOffset)
	}

	for i := 0; i < p.params.PartialRounds; i++ {
		p.applyPartialRound(&keysOffset)
	}

	for i := 0; i < p.params.FullRoundsEnd; i++ {
		p.applyFullRound(&keysOffset)
	}

	return p.input[1].Bytes()
}

func (p *Poseidon) applyFullRound(keysOffset *int) {
	// Add current round constant to all elements of input
	for i := 0; i < len(p.input); i++ {
		p.input[i] = *p.input[i].Add(&p.input[i], &p.params.RoundKeys[*keysOffset])
		*keysOffset++
	}

	// Apply quintic SBox to every element
	for i := 0; i < len(p.input); i++ {
		QuinticSbox(&p.input[i])
	}

	p.input = *mulVec(&p.params.MDSMatrix, &p.input)
}

func (p *Poseidon) applyPartialRound(keysOffset *int) {
	// Add current round constant to all elements of input
	for i := 0; i < len(p.input); i++ {
		p.input[i] = *p.input[i].Add(&p.input[i], &p.params.RoundKeys[*keysOffset])
		*keysOffset++
	}

	// Apply quintic SBox to the first element
	QuinticSbox(&p.input[0])

	p.input = *mulVec(&p.params.MDSMatrix, &p.input)
}

func mulVec(a *[][]ristretto.Scalar, b *[]ristretto.Scalar) *[]ristretto.Scalar {
	result := make([]ristretto.Scalar, len(*b))

	for j := 0; j < len(*a); j++ {
		line := make([]ristretto.Scalar, len(*b))

		for k := 0; k < len((*a)[j]); k++ {
			line[k].Mul(&(*a)[j][k], &(*b)[k])
		}

		for k := 0; k < len(line); k++ {
			result[j].Add(&result[j], &line[k])
		}
	}

	return &result
}

// QuinticSbox will set *a to a^5
func QuinticSbox(a *ristretto.Scalar) {
	c := *a
	for k := 0; k < 4; k++ {
		a.Mul(a, &c)
	}
}

// Pad will fill the input with zeroed scalars until its length equal the parametrization width
func (p *Poseidon) Pad() {
	dif := p.params.Width - len(p.input)
	if dif > 0 {
		pad := make([]ristretto.Scalar, dif)
		p.input = append(p.input, pad...)
	}
}

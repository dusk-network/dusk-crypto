package poseidon

import (
	"bytes"
	"encoding/binary"
	"errors"

	ristretto "github.com/bwesterb/go-ristretto"
)

// Params represents the parameters for the hash calculation
type Params struct {
	Width               int
	FullRoundsBeginning int
	PartialRounds       int
	FullRoundsEnd       int
	RoundKeys           []ristretto.Scalar
	MDSMatrix           [][]ristretto.Scalar
}

// Default will generate a default parametrization for Poseidon
func Default() Params {
	return Params{
		Width:               5,
		FullRoundsBeginning: 4,
		PartialRounds:       59,
		FullRoundsEnd:       4,
		RoundKeys:           DefaultRoundKeys[:],
		MDSMatrix:           DefaultMDSMatrix,
	}
}

// GenerateMDSMatrix will create a width x width MDS matrix
func GenerateMDSMatrix(width int) (*[][]ristretto.Scalar, error) {
	// Initialize mds
	mds := make([][]ristretto.Scalar, width)
	for i := 0; i < width; i++ {
		mds[i] = make([]ristretto.Scalar, width)
	}

	xs := make([]ristretto.Scalar, width)
	ys := make([]ristretto.Scalar, width)

	for i := 0; i < width; i++ {
		x, err := UnsignedNumberToScalar(i)
		if err != nil {
			return nil, err
		}
		y, err := UnsignedNumberToScalar(i + width)
		if err != nil {
			return nil, err
		}

		xs[i] = *x
		ys[i] = *y
	}

	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			mds[i][j].Add(&xs[i], &ys[j])
			mds[i][j].Inverse(&mds[i][j])
		}
	}

	return &mds, nil
}

// UnsignedNumberToScalar will attempt to convert an unsigned number to a Scalar
func UnsignedNumberToScalar(number interface{}) (*ristretto.Scalar, error) {
	switch v := number.(type) {

	case uint8, uint16, uint, uint32, uint64:
		buf := bytes.NewBuffer([]byte{})
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			return nil, err
		}

		b := buf.Bytes()

		// Pad b with zeroes
		for i := 0; i < 2; i++ {
			if len(b) < 4*(i+1) {
				l := 4*(i+1) - len(b)
				t := make([]byte, l)
				b = append(b, t...)
			}
		}

		b1 := make([]byte, 8)
		for i := 0; i < 4; i++ {
			b1[i] = b[i]
		}
		s1, n := binary.Uvarint(b1[:])
		if n == 0 {
			return nil, errors.New("Buffer is too small")
		}

		b2 := make([]byte, 8)
		for i := 0; i < 4; i++ {
			b2[i] = b[i+4]
		}
		s2, n := binary.Uvarint(b2[:])
		if n == 0 {
			return nil, errors.New("Buffer is too small")
		}

		s := ristretto.Scalar{uint32(s1), uint32(s2), 0, 0, 0, 0, 0, 0}
		return &s, nil

	default:
		return nil, errors.New("Conversion allowed only for unsigned numeric types")

	}
}

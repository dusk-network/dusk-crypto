package merkletree

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/dusk-network/dusk-crypto/poseidon"
)

type TestPayload struct {
	x string
}

func (t TestPayload) CalculateHash() ([]byte, error) {
	p := poseidon.New()

	_, err := p.Write([]byte(t.x))
	if err != nil {
		return nil, err
	}

	return p.Sum(nil), nil
}

type Fixture struct {
	Payloads     []Payload
	ExpectedHash []byte
}

var testTable = []Fixture{
	{
		Payloads: []Payload{
			TestPayload{
				x: "Bella",
			},
			TestPayload{
				x: "Ciao",
			},
			TestPayload{
				x: "Ndo",
			},
			TestPayload{
				x: "Scappi",
			},
		},
		ExpectedHash: []byte{
			7, 47, 154, 13, 74, 201, 26, 124, 103, 145, 133, 192, 99, 161, 206, 218, 129, 125, 239, 96, 35, 213, 84, 53, 77, 200, 187, 170, 247, 223, 195, 9,
		},
	},

	{
		Payloads: []Payload{
			TestPayload{
				x: "Hello",
			},
			TestPayload{
				x: "Hi",
			},
			TestPayload{
				x: "Hey",
			},
			TestPayload{
				x: "Hola",
			},
		},
		ExpectedHash: []byte{
			130, 42, 56, 141, 97, 229, 176, 49, 37, 250, 172, 139, 222, 90, 162, 118, 196, 34, 212, 82, 6, 110, 54, 148, 75, 225, 100, 100, 53, 3, 209, 4,
		},
	},
}

type uTest func(t *Tree, fixture Fixture, index int) error

func injectTest(t *testing.T, u uTest) {
	for i := 0; i < len(testTable); i++ {
		currentFixture := testTable[i]

		tree, err := NewTree(testTable[i].Payloads)
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		if err := u(tree, currentFixture, i); err != nil {
			t.Errorf(err.Error())
		}

	}
}

var newTreeTest = func(tree *Tree, f Fixture, i int) error {
	eh := f.ExpectedHash
	root := tree.MerkleRoot

	if bytes.Compare(root, eh) != 0 {
		return fmt.Errorf("expecting hash %v but got %v", eh, root)
	}
	return nil
}

var merkleRootTest = func(tree *Tree, f Fixture, i int) error {
	root := tree.MerkleRoot
	if bytes.Compare(root, f.ExpectedHash) != 0 {
		return fmt.Errorf("expecting hash %v but got %v", f.ExpectedHash, root)
	}
	return nil
}

var rebuildTreeTest = func(tree *Tree, f Fixture, i int) error {
	root := tree.MerkleRoot
	if err := tree.RebuildTree(); err != nil {
		return err
	}

	if bytes.Compare(root, f.ExpectedHash) != 0 {
		return fmt.Errorf("expecting hash %v but got %v", f.ExpectedHash, root)
	}
	return nil
}

var verifyTree = func(tree *Tree, f Fixture, i int) error {

	v1, err := VerifyTree(tree)
	if err != nil {
		return err
	}

	if !v1 {
		return errors.New("expected tree to be valid")
	}

	tree.Root.Hash = []byte{1}
	tree.MerkleRoot = []byte{1}

	v2, err := VerifyTree(tree)
	if err != nil {
		return err
	}

	if v2 {
		return errors.New("expected tree to be invalid")
	}
	return nil
}

var verifyContent = func(tree *Tree, f Fixture, i int) error {

	for _, payload := range f.Payloads {
		v, err := tree.VerifyContent(payload)
		if err != nil {
			return err
		}
		if !v {
			return fmt.Errorf("encountered invalid payload %s", payload)
		}
	}

	v, err := tree.VerifyContent(TestPayload{x: "NotInTestTable"})
	if err != nil {
		return err
	}

	if v {
		return errors.New("verification should have failed")
	}

	return nil
}

func TestMerkleTree_RebuildWithPayload(t *testing.T) {
	for i := 0; i < len(testTable)-1; i++ {
		tree, err := NewTree(testTable[i].Payloads)
		if err != nil {
			t.Error("unexpected error:  ", err)
		}
		err = tree.RebuildTreeUsing(testTable[i+1].Payloads)
		if err != nil {
			t.Error("unexpected error:  ", err)
		}
		if bytes.Compare(tree.MerkleRoot, testTable[i+1].ExpectedHash) != 0 {
			t.Errorf("expected hash equal to %v got %v", testTable[i+1].ExpectedHash, tree.MerkleRoot)
		}
	}
}

func TestSuite(t *testing.T) {
	injectTest(t, newTreeTest)
	injectTest(t, merkleRootTest)
	injectTest(t, rebuildTreeTest)
	t.Run("Rebuilding the MerkleTree with the payload yields the same MerkleTree", TestMerkleTree_RebuildWithPayload)
	injectTest(t, verifyTree)
	injectTest(t, verifyContent)
}

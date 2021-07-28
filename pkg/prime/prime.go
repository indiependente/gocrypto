package prime

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	modLen = 20
)

// Generate returns a prime of bits length and panics on error.
func Generate(bits int) *big.Int {
	p, err := rand.Prime(rand.Reader, bits+1)
	if err != nil {
		panic(err)
	}
	return p
}

type Zp struct {
	P *big.Int
}

func NewZp(bitlen int) Zp {
	return Zp{
		P: Generate(bitlen),
	}
}

func (zp Zp) RandomZp() (*big.Int, error) {
	n, err := rand.Int(rand.Reader, zp.P)
	if err != nil {
		return nil, fmt.Errorf("could not generate integer in Zp: %w", err)
	}
	return n, nil
}

func (zp Zp) RandomMultiplicativeZp() (*big.Int, error) {
	var (
		n   *big.Int
		err error
	)
	for {
		n, err = rand.Int(rand.Reader, zp.P)
		if err != nil {
			return nil, fmt.Errorf("could not generate integer in multiplicative Zp: %w", err)
		}
		if n.Cmp(big.NewInt(0)) > 0 {
			break
		}
	}

	return n, nil
}

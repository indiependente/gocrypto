package secretsharing

import (
	"fmt"
	"math/big"

	"github.com/indiependente/gocrypto/pkg/horner"
	"github.com/indiependente/gocrypto/pkg/prime"
)

const (
	// ErrNotEnoughShares is returned on an attempt to retrieve the secret with not enough shares.
	ErrNotEnoughShares Error = "not enough shares provided"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

// SecretSharing implements Shamir's Secret Sharing.
type SecretSharing struct {
	// zp Defines Zp
	zp prime.Zp
	// coeff polynomial coefficients
	coeff []*big.Int
	// shares to give to participants
	shares []*Share
	// ids participants identities
	ids []*big.Int

	// number of shares
	n int
	// threshold
	k int
}

func (ss SecretSharing) Shares() []*Share {
	return ss.shares[:]
}

func New(secret []byte, k, n int) (*SecretSharing, error) {
	shares := make([]*Share, k+1)
	coeff := make([]*big.Int, k)
	ids := make([]*big.Int, n+1)

	s := new(big.Int)
	s.SetBytes(secret)
	zp := prime.NewZp(s.BitLen())
	s.Mod(s, zp.P)

	shares[0] = NewShare(0, s)
	coeff[0] = s

	for i := 0; i <= n; i++ {
		ids[i] = big.NewInt(int64(i))
	}

	for i := 1; i < k; i++ {
		c, err := zp.RandomZp()
		if err != nil {
			return nil, fmt.Errorf("could not generate coefficient %d: %w", i, err)
		}
		coeff[i] = c
	}

	return &SecretSharing{
		zp:     zp,
		shares: shares,
		coeff:  coeff,
		ids:    ids,
		k:      k,
		n:      n,
	}, nil
}

func (ss *SecretSharing) ComputeShares() {
	for i := 1; i <= ss.k; i++ {
		ss.shares[i] = NewShare(ss.ids[i].Int64(), horner.PolyEval(ss.coeff, ss.ids[i], ss.zp.P))
	}
}

func (ss SecretSharing) ComputeSecret(shares []*Share) ([]byte, error) {
	if len(shares) <= ss.k {
		return nil, ErrNotEnoughShares
	}
	secret := big.NewInt(0)

	for j := 1; j <= ss.k; j++ {
		prod := big.NewInt(1)
		idJ := big.NewInt(shares[j].id.Int64())
		yJ := big.NewInt(shares[j].value.Int64())

		for t := 1; t <= ss.k; t++ {
			if t == j {
				continue
			}
			idT := big.NewInt(shares[t].id.Int64())

			var sub big.Int
			sub.Sub(ss.zp.P, idJ)
			sub.Add(&sub, idT)
			sub.ModInverse(&sub, ss.zp.P)
			sub.Mul(&sub, idT)
			sub.Mod(&sub, ss.zp.P)

			prod.Mul(prod, &sub).Mod(prod, ss.zp.P)
		}

		prod.Mul(prod, yJ).Mod(prod, ss.zp.P)

		secret.Add(secret, prod).Mod(secret, ss.zp.P)

	}

	return secret.Bytes(), nil
}

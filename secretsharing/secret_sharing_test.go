package secretsharing

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/indiependente/gocrypto/pkg/prime"
	"github.com/stretchr/testify/assert"
)

func TestSecretSharing_ComputeSecret(t *testing.T) {
	t.Parallel()
	type fields struct {
		zp     prime.Zp
		coeff  []*big.Int
		shares []*Share
		ids    []*big.Int
		n      int
		k      int
	}

	tests := []struct {
		name   string
		fields fields
		secret []byte
		want   []byte
	}{
		{
			name: "S=12_k=3_n=5_shares_given",
			fields: fields{
				zp: prime.Zp{
					P: big.NewInt(19),
				},
				coeff: []*big.Int{
					big.NewInt(12),
					big.NewInt(11),
					big.NewInt(2),
				},
				shares: []*Share{
					NewShare(0, big.NewInt(12)),
					NewShare(1, big.NewInt(6)),
					NewShare(2, big.NewInt(4)),
					NewShare(3, big.NewInt(6)),
				},
				ids: []*big.Int{
					big.NewInt(0),
					big.NewInt(1),
					big.NewInt(2),
					big.NewInt(3),
					big.NewInt(4),
					big.NewInt(5),
				},
				n: 5,
				k: 3,
			},

			want: big.NewInt(12).Bytes(),
		},
		{
			name:   "build shares - secret number",
			secret: big.NewInt(13).Bytes(),
			fields: fields{
				k: 3,
				n: 5,
			},
			want: big.NewInt(13).Bytes(),
		},
		{
			name:   "build shares - secret string",
			secret: []byte("the secret"),
			fields: fields{
				k: 3,
				n: 5,
			},
			want: []byte("the secret"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var (
				ss  *SecretSharing
				err error
			)
			if len(tt.secret) == 0 {
				ss = &SecretSharing{
					zp:     tt.fields.zp,
					coeff:  tt.fields.coeff,
					shares: tt.fields.shares,
					ids:    tt.fields.ids,
					n:      tt.fields.n,
					k:      tt.fields.k,
				}
			} else {
				ss, err = New(tt.secret, tt.fields.k, tt.fields.n)
				assert.NoError(t, err)
				ss.ComputeShares()
				shares := ss.Shares()
				rand.Seed(time.Now().Unix())
				idxs := rand.Perm(ss.k)
				for _, i := range idxs {
					tt.fields.shares = append(tt.fields.shares, shares[i])
				}
			}

			got, err := ss.ComputeSecret(tt.fields.shares)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

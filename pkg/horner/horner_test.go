package horner

import (
	"math/big"
	"testing"

	"github.com/indiependente/gocrypto/pkg/prime"
	"github.com/stretchr/testify/assert"
)

func TestPolyEval(t *testing.T) {
	t.Parallel()
	type args struct {
		coefficients []*big.Int
		x            *big.Int
		p            *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Evaluate x = 0",
			args: args{
				coefficients: []*big.Int{
					big.NewInt(12),
					big.NewInt(11),
					big.NewInt(2),
				},
				p: prime.Generate(1024),
				x: big.NewInt(0),
			},
			want: big.NewInt(12),
		},
		{
			name: "Evaluate x = 1",
			args: args{
				coefficients: []*big.Int{
					big.NewInt(12),
					big.NewInt(11),
					big.NewInt(2),
				},
				p: prime.Generate(1024),
				x: big.NewInt(1),
			},
			want: big.NewInt(25),
		},
		{
			name: "Evaluate x = 2",
			args: args{
				coefficients: []*big.Int{
					big.NewInt(12),
					big.NewInt(11),
					big.NewInt(2),
				},
				x: big.NewInt(2),
				p: prime.Generate(1024),
			},
			want: big.NewInt(42),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			y := PolyEval(tt.args.coefficients, tt.args.x, tt.args.p)
			assert.Equal(t, tt.want, y)
		})
	}
}

package secretsharing

import (
	"encoding/base64"
	"fmt"
	"math/big"
)

type Share struct {
	id    *big.Int
	value *big.Int
}

func NewShare(id int64, v *big.Int) *Share {
	return &Share{
		id:    big.NewInt(id),
		value: v,
	}
}

func (s Share) String() string {
	return fmt.Sprintf("ID: %d\tValue: %d\tBase64: %s\n", s.id.Int64(), s.value.Int64(), s.Base64())
}

func (s Share) Base64() string {
	return base64.StdEncoding.EncodeToString(s.value.Bytes())
}

package verification

import (
	"math/big"

	Acc "../Acc"
)

func Verify(u, W, Accumulator big.Int, key Acc.Rsa_key) bool {
	e := Acc.Hprime(u)
	N := key.N
	if W.Exp(&W, &e, &N).Cmp(&Accumulator) == 0 {
		return true
	}

}

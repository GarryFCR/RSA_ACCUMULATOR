package verification

import (
	"math/big"

	"../Acc"
)

//Verification is simply
//W^e (mod N) == Acc
func Verify(u, W, Accumulator, N big.Int) bool {
	e := Acc.Hprime(u)
	Acc_dash := new(big.Int).Exp(&W, &e, &N)
	return Acc_dash.Cmp(&Accumulator) == 0

}

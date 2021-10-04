package verification

import (
	"math/big"

	"../Acc"
)

//Verification is simply
//W^e (mod N) == Acc
func Verify(args []big.Int) bool {

	u, W, Accumulator, N := args[0], args[1], args[2], args[3]
	e := Acc.Hprime(u)
	Acc_dash := new(big.Int).Exp(&W, &e, &N)
	return Acc_dash.Cmp(&Accumulator) == 0

}

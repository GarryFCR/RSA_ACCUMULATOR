package verification

import (
	"math/big"

	Acc "../Acc"
)

//Verification is simply
//W^e (mod N) == Acc
func Verify(u, W, Accumulator, N big.Int) bool {

	e := Acc.Hprime(u)
	Acc_dash := new(big.Int).Exp(&W, &e, &N)

	//fmt.Println("Acc calculated:", Acc_dash, "e:", e, "N:", N)

	return Acc_dash.Cmp(&Accumulator) == 0

}

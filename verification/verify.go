package verification

import (
	"math/big"

	"github.com/GarryFCR/RSA_ACCUMULATOR/Rsa_accumulator"
)

//Verification is simply
//W^e (mod N) == Acc
func Verify(args []big.Int) bool {

	u, W, Accumulator, N := args[0], args[1], args[2], args[3]
	e := Rsa_accumulator.Hprime(u)
	Acc_dash := new(big.Int).Exp(&W, &e, &N)
	return Acc_dash.Cmp(&Accumulator) == 0

}

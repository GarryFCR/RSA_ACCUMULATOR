package witness

import (
	"math/big"

	Acc "./github.com/GarryFCR/RSA_ACCUMULATOR/src/Acc"
)

func Generate_witness(Accumulator, u big.Int, key Acc.Rsa_key) big.Int {

	e := Acc.Hprime(u)
	N := key.N
	q := key.Q
	p := key.P

	phi := new(big.Int).Mul(new(big.Int).Sub(&q, big.NewInt(1)), new(big.Int).Sub(&p, big.NewInt(1)))

	inverse := new(big.Int).ModInverse(&e, phi)
	W := new(big.Int).Exp(&Accumulator, inverse, &N)
	return *W
}

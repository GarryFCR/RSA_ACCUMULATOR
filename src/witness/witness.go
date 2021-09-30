package witness

import (
	"math/big"

	Acc "../Acc"
)

//Generation of witness is multiplication of all primes mapped from members except the one we
//are proving,prod(say) then,
// Witness = G^prod(mod N)
func Generate_witness(u big.Int, key Acc.Rsa_key, U []big.Int) big.Int {

	N := key.N

	Primes := make([]big.Int, len(U))
	G := key.G
	for i, u_dash := range U {
		Primes[i] = Acc.Hprime(u_dash)
		if u_dash.Cmp(&u) != 0 {

			G.Exp(&G, &Primes[i], &N)
		}

	}

	return G
}

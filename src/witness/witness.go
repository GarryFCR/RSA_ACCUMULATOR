package witness

import (
	"math/big"

	Acc "../Acc"
)

type Witness_list struct {
	Acc  big.Int
	List map[string]big.Int
}

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

func (witness *Witness_list) Precompute_witness(G_prev big.Int, U []big.Int, accumulator *Acc.Rsa_Acc) {

	if len(U) == 1 {
		u := U[0]
		witness.List[u.String()] = G_prev
		witness.Acc = accumulator.Acc
		return
	}

	A := U[:len(U)/2]
	B := U[len(U)/2:]
	G1 := G_prev
	G2 := G_prev
	N := accumulator.N

	for _, u := range B {

		e1 := Acc.Hprime(u)
		G1.Exp(&G1, &e1, &N)

	}

	for _, w := range A {
		e2 := Acc.Hprime(w)
		G2.Exp(&G2, &e2, &N)
	}
	//fmt.Println("A,B:", G1, G2)
	witness.Precompute_witness(G1, A, accumulator)
	witness.Precompute_witness(G2, B, accumulator)

}

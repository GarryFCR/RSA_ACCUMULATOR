package Acc

import (
	"math/big"
)

type Witness_list struct {
	Acc  big.Int
	List map[string]big.Int
}

//Whenever the set is passed or it changes there is a computation of new witnesses
func (witness *Witness_list) Precompute_witness(G_prev big.Int, U []big.Int, accumulator *Rsa_Acc) {

	if len(U) == 1 {
		witness.List[U[0].String()] = G_prev
		witness.Acc = accumulator.Acc
		return
	}

	A := U[:len(U)/2]
	B := U[len(U)/2:]
	G1 := G_prev
	G2 := G_prev
	N := accumulator.N

	for _, u := range B {
		e1 := Hprime(u)

		G1.Exp(&G1, &e1, &N)
	}

	for _, w := range A {
		e2 := Hprime(w)
		G2.Exp(&G2, &e2, &N)
	}
	//fmt.Println(G1, G2)
	witness.Precompute_witness(G1, A, accumulator)
	witness.Precompute_witness(G2, B, accumulator)
}

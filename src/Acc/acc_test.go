package Acc

import (
	"crypto/rand"
	//"fmt"
	"math/big"
	"testing"
)

func TestGenerate_acc(t *testing.T) {
	//hash2prime-----------------------------------------------------------------------
	H, _ := rand.Int(rand.Reader, big.NewInt(512))
	E := Hprime(*H)

	if !E.ProbablyPrime(10) {
		t.Fatalf("Not a prime")
	}
	//generate_acc---------------------------------------------------------------------
	key := Rsa_keygen(int(32))
	if key.G.Cmp(&key.N) == 1 && key.N.BitLen() == 32 {
		t.Fatalf("Incorrect generator")
	}

	U := []big.Int{*big.NewInt(123), *big.NewInt(124), *big.NewInt(125), *big.NewInt(126)}
	Accumulator := Generate_Acc(key, U)
	G := key.G
	for _, u := range U {
		x := Hprime(u)
		G.Exp(&G, &x, &key.N)
	}
	if G.Cmp(&Accumulator.Acc) != 0 {
		t.Fatalf("Incorrect accumulator")
	}

	//witness---------------------------------------------------------------------------
	list := make(map[string]big.Int, len(Accumulator.U))
	w := &Witness_list{Acc: Accumulator.Acc, List: list}
	w.Precompute_witness(Accumulator.G, Accumulator.U, Accumulator)

	witnesses := make([]big.Int, len(Accumulator.U))
	for i, u := range Accumulator.U {
		witnesses[i] = generate_witness(u, key, Accumulator.U)
	}

	for i, u := range U {
		witness := w.List[u.String()]
		if witnesses[i].Cmp(&witness) != 0 {
			t.Fatalf("Incorrect witness")
		}
	}

}

//Generation of witness is multiplication of all primes mapped from members except the one we
//are proving,prod(say) then,
// Witness = G^prod(mod N)
func generate_witness(u big.Int, key Rsa_key, U []big.Int) big.Int {

	N := key.N

	Primes := make([]big.Int, len(U))
	G := key.G
	for i, u_dash := range U {
		Primes[i] = Hprime(u_dash)
		if u_dash.Cmp(&u) != 0 {

			G.Exp(&G, &Primes[i], &N)
		}

	}

	return G
}

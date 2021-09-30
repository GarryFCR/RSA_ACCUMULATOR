package Acc

import (
	"crypto/rand"
	"crypto/rsa"
	"math/big"
)

type Rsa_key struct {
	N big.Int //Z/nZ
	G big.Int // Generator

}
type Rsa_Acc struct {
	Acc big.Int   // Accumulator
	U   []big.Int // set of members
	N   big.Int   //Z/nZ
	G   big.Int   // Generator

}

func Rsa_keygen(lambda int) Rsa_key {

	pk, _ := rsa.GenerateKey(rand.Reader, lambda)
	var F *big.Int
	N := pk.PublicKey.N

	for {
		F, _ = rand.Int(rand.Reader, N)
		if new(big.Int).Mod(F, pk.Primes[0]).Cmp(big.NewInt(0)) != 0 &&
			new(big.Int).Mod(F, pk.Primes[1]).Cmp(big.NewInt(0)) != 0 &&
			F != big.NewInt(1) {
			break
		}
	}

	G := new(big.Int).Exp(F, big.NewInt(2), pk.PublicKey.N)

	return Rsa_key{
		N: *N,
		G: *G,
	}
}

func Generate_Acc(key Rsa_key, U []big.Int) *Rsa_Acc {

	Primes := make([]big.Int, len(U))
	G := key.G
	for i, u := range U {
		Primes[i] = Hprime(u)
		G.Exp(&G, &Primes[i], &key.N)
	}
	//fmt.Println(Primes)

	return &Rsa_Acc{
		Acc: G,
		U:   U,
		N:   key.N,
		G:   key.G,
	}

}

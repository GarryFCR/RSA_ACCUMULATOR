package Rsa_accumulator

import (
	"crypto/rand"
	"crypto/rsa"
	"math/big"
)

//generated by the verifier
type Rsa_key struct {
	N big.Int //Z/nZ
	G big.Int // Generator

}

//for the user to store
type Rsa_Acc struct {
	Acc big.Int   // Accumulator
	U   []big.Int // set of members
	N   big.Int   //Z/nZ
	G   big.Int   // Generator

}

//Generate N and G
//lamda is the bit size of primes
func Rsa_keygen(lambda int) Rsa_key {

	pk, _ := rsa.GenerateKey(rand.Reader, lambda)
	var F *big.Int
	var err error
	N := pk.PublicKey.N

	for {
		F, err = rand.Int(rand.Reader, N)
		if new(big.Int).GCD(nil, nil, F, N).Cmp(big.NewInt(1)) == 0 && err == nil {
			break
		}
	}
	G := new(big.Int).Exp(F, big.NewInt(2), pk.PublicKey.N)
	return Rsa_key{
		N: *N,
		G: *G,
	}
}

//Generate the accumulator
func Generate_Acc(key Rsa_key, U []big.Int) *Rsa_Acc {

	Primes := make([]big.Int, len(U))
	G := key.G

	for i, u := range U {
		Primes[i] = Hprime(u)
		G.Exp(&G, &Primes[i], &key.N)
	}

	return &Rsa_Acc{
		Acc: G,
		U:   U,
		N:   key.N,
		G:   key.G,
	}

}
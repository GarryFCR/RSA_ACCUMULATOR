package Acc

import (
	"crypto/rand"
	"crypto/rsa"
	"math/big"
)

type rsa_key struct {
	N big.Int //Z/nZ
	G big.Int // Generator
	p big.Int //prime 1
	q big.Int //prime 2

}

func Rsa_keygen(lambda int) rsa_key {

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

	return rsa_key{
		N: *N,
		G: *G,
		p: *pk.Primes[0],
		q: *pk.Primes[1],
	}
}

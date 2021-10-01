package Acc

import (
	"math/big"
)

func (c *Rsa_Acc) Add_member(u big.Int) {

	e := Hprime(u)
	newAcc := new(big.Int).Exp(&c.Acc, &e, &c.N)

	newSet := append(c.U[:], u)
	c.Acc = *newAcc
	c.U = newSet

}

func (c *Rsa_Acc) Delete_member(u big.Int) {

	var NewSet []big.Int
	var i int
	for i = 0; i < len(c.U); i++ {
		if c.U[i].Cmp(&u) == 0 {
			NewSet = append(c.U[:i], c.U[i+1:]...)
			break
		}
	}

	key := Rsa_key{N: c.N, G: c.G}
	NewAcc := Generate_Acc(key, NewSet)
	c.Acc = NewAcc.Acc
	c.U = NewSet

}

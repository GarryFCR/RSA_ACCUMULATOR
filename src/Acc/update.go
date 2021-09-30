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

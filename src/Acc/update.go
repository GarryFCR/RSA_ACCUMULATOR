package Acc

import (
	"math/big"
)

func (c *Rsa_Acc) Add_member(u big.Int, w *Witness_list) {

	e := Hprime(u)
	newAcc := new(big.Int).Exp(&c.Acc, &e, &c.N)

	newSet := append(c.U[:], u)
	c.Acc = *newAcc
	c.U = newSet
	w.Precompute_witness(c.G, c.U, c)

}

func (c *Rsa_Acc) Delete_member(u big.Int, w *Witness_list) {

	var newSet []big.Int
	var i int
	for i = 0; i < len(c.U); i++ {
		if c.U[i].Cmp(&u) == 0 {
			newSet = append(c.U[:i], c.U[i+1:]...)
			break
		}
	}

	newAcc := w.List[u.String()]
	c.Acc = newAcc
	c.U = newSet
	list := make(map[string]big.Int, len(c.U))
	w.List = list
	w.Precompute_witness(c.G, c.U, c)

}

package main

import (
	"fmt"
	"math/big"

	"./Acc"
	verify "./verification"
)

func main() {

	//Verifier on-chain------------------------------------------------------------------------
	//generate N for the RSA group
	//generate G from RSA group
	//key is then sent to user
	key := Acc.Rsa_keygen(32)
	fmt.Println("Public Key:", &key)

	//User off-chain----------------------------------------------------------------------------
	//Example set
	U := []big.Int{*big.NewInt(123), *big.NewInt(124), *big.NewInt(125), *big.NewInt(126)}

	//Generate the accumulator for the above set
	//send Acc to verifier
	Accumulator := Acc.Generate_Acc(key, U)
	fmt.Println("Accumulator:", Accumulator)

	//Initialising of witness
	list1 := make(map[string]big.Int, len(Accumulator.U))
	w := &Acc.Witness_list{Acc: Accumulator.Acc, List: list1}

	//Precompute witness--------------------------------------------------------------------------
	w.Precompute_witness(Accumulator.G, Accumulator.U, Accumulator)

	//Adding a member
	Accumulator.Add_member(*big.NewInt(127), w)
	fmt.Println("witness map after adding 127", w.List)

	//Deleting a member
	Accumulator.Delete_member(*big.NewInt(126), w)
	fmt.Println("witness map after removing 126", w.List)

	//Witness for a member
	//send (witness,member) to verifier
	W1 := w.List[Accumulator.U[2].String()]

	//Verifier on-chain------------------------------------------------------------------------
	//Verification
	if verify.Verify(*big.NewInt(125), W1, Accumulator.Acc, key.N) {
		fmt.Printf("%v is a valid member\n", big.NewInt(125))
	} else {
		fmt.Printf("%v is not a member\n", big.NewInt(125))
	}

}

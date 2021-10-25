package main

import (
	"fmt"
	"math/big"

	"github.com/man2706kum/RSA_ACCUMULATOR/Acc"
	verify "github.com/man2706kum/RSA_ACCUMULATOR/verification"
)

func main() {

	//Verifier on-chain------------------------------------------------------------------------
	//generate N for the RSA group
	//generate G from RSA group
	//key is then sent to user
	key := Acc.Rsa_keygen(32)
	fmt.Println("Public Key (N, G):", &key)

	//User off-chain----------------------------------------------------------------------------
	//Example set
	U := []big.Int{*big.NewInt(123), *big.NewInt(124), *big.NewInt(125), *big.NewInt(126)}

	//Generate the accumulator for the above set
	//send Acc to verifier
	Accumulator := Acc.Generate_Acc(key, U)
	fmt.Println("\n\nAccumulator(Acc, U, N, G):", Accumulator)

	//Initialising of witness
	w := Accumulator.Witness_int()

	//Precompute witness
	w.Precompute_witness(Accumulator.G, Accumulator.U, Accumulator)

	//Adding a member
	Accumulator.Add_member(*big.NewInt(127), w)
	fmt.Println("\n\nwitness map after adding 127: ", w.List)

	//Deleting a member
	Accumulator.Delete_member(*big.NewInt(126), w)
	fmt.Println("\n\nwitness map after removing 126: ", w.List)

	//Witness for a member
	//send (witness,member) to verifier
	W1 := w.List[Accumulator.U[2].String()]

	//Verifier on-chain------------------------------------------------------------------------
	//Verification
	args := []big.Int{*big.NewInt(125), W1, Accumulator.Acc, key.N}
	fmt.Println("\n\nChecking if 125 is a valid member...")
	if verify.Verify(args) {
		fmt.Printf("%v is a valid member\n", big.NewInt(125))
	} else {
		fmt.Printf("%v is not a member\n", big.NewInt(125))
	}

}

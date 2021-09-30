package main

import (
	"fmt"
	"math/big"

	"./Acc"
	verify "./verification"
	"./witness"
)

func main() {

	//Generation of a Hidden group order
	key := Acc.Rsa_keygen(12)
	fmt.Println("Public Key:", &key)

	//Example set
	U := []big.Int{*big.NewInt(123), *big.NewInt(124), *big.NewInt(125), *big.NewInt(126)}

	//Generate the accumulator for th above set
	Accumulator := Acc.Generate_Acc(key, U)
	fmt.Println("Acc:", Accumulator)

	//witness of a member
	W1 := witness.Generate_witness(*big.NewInt(125), key, U)

	//Verification
	if verify.Verify(*big.NewInt(125), W1, Accumulator.Acc, key.N) {
		fmt.Printf("%v is a valid member\n", big.NewInt(125))
	} else {
		fmt.Printf("%v is not a member\n", big.NewInt(125))
	}

	//witness of a non-member
	W2 := witness.Generate_witness(*big.NewInt(127), key, U)

	if verify.Verify(*big.NewInt(15), W2, Accumulator.Acc, key.N) {
		fmt.Printf("%v is a valid member\n", big.NewInt(127))
	} else {
		fmt.Printf("%v is not a member\n", big.NewInt(127))
	}

	Accumulator.Add_member(*big.NewInt(127))
	fmt.Println("Acc:", Accumulator)

}

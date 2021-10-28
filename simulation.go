package main

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/GarryFCR/RSA_ACCUMULATOR/Acc"
	verify "github.com/GarryFCR/RSA_ACCUMULATOR/verification"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	loops       = 5
	sliceLength = 100
)

var (
	source  = rand.NewSource(time.Now().UnixNano())
	key     = Acc.Rsa_keygen(512) //group order should be greater than the elements
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	random  = rand.New(source)
)

func main() {

	fmt.Println("For set size:", sliceLength)
	for loop := 0; loop < loops; loop++ {
		array := make([]big.Int, sliceLength)
		for i, _ := range array {
			array[i] = *big.NewInt(0).SetBytes(randomAddress().Bytes())
		}
		start1 := time.Now()
		Accumulator := Acc.Generate_Acc(key, array)
		w := Accumulator.Witness_int()
		elapsed1 := time.Since(start1)
		log.Printf("Accumulator generation time :%s", elapsed1)

		start2 := time.Now()
		w.Precompute_witness(Accumulator.G, Accumulator.U, Accumulator)
		elapsed2 := time.Since(start2)
		log.Printf("witness computation time :%s", elapsed2)

		t, fp, fn := check(array, w, Accumulator)
		fmt.Println("loop:", loop, "t:", t, "fp:", fp, "fn:", fn)
		fmt.Println()
	}

}

func randomAddress() sdk.AccAddress {
	n := 20
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[random.Intn(len(letters))]
	}
	return auth.NewModuleAddress(string(b))
}

func check(original []big.Int, w *Acc.Witness_list, Accumulator *Acc.Rsa_Acc) (int, int, int) {
	t := 0
	fp := 0
	fn := 0
	for _, c := range original {
		args := []big.Int{c, w.List[c.String()], Accumulator.Acc, key.N}
		if verify.Verify(args) {
			t++
		} else {
			fn++
		}
	}
	for i := 0; i < sliceLength; i++ {
		a := *big.NewInt(0).SetBytes(randomAddress().Bytes())
		args := []big.Int{a, w.List[a.String()], Accumulator.Acc, key.N}
		if verify.Verify(args) {
			fp++
		}
	}
	return t, fp, fn
}

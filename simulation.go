package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/GarryFCR/RSA_ACCUMULATOR/Acc"
	verify "github.com/GarryFCR/RSA_ACCUMULATOR/verification"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	loops       = 10
	sliceLength = 1000
)

var (
	source  = rand.NewSource(time.Now().UnixNano())
	key     = Acc.Rsa_keygen(128) //group order should be greater than the elements
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	random  = rand.New(source)
)

func main() {

	a := big.NewInt(0).SetBytes(randomAddress().Bytes())
	fmt.Println(a.Div(a, big.NewInt(9223372036854775807)).String())

	for loop := 0; loop < loops; loop++ {
		array := make([]big.Int, sliceLength)
		for i, _ := range array {
			array[i] = *big.NewInt(0).SetBytes(randomAddress().Bytes())
		}
		Accumulator := Acc.Generate_Acc(key, array)
		w := Accumulator.Witness_int()
		w.Precompute_witness(Accumulator.G, Accumulator.U, Accumulator)
		t, fp, fn := check(array, w, Accumulator)
		fmt.Println("loop:", loop, "t:", t, "fp:", fp, "fn:", fn)
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

package main

import (
	"fmt"

	"github.com/givek/crypto-pals-challenges/set1"
)

func main() {

	// hexStr := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	// fmt.Println(set1.HexToBase64(hexStr))

	x := []byte("1c0111001f010100061a024b53535009181c")

	y := []byte("686974207468652062756c6c277320657965")

	fmt.Println(set1.XorBase16(x, y))
}

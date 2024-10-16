package set1

import (
	"strconv"
)

func binToSextet(bin [6]byte) int {

	v := [6]byte{32, 16, 8, 4, 2, 1}

	res := 0

	for i, b := range bin {

		res = res + int(b*v[i])

	}

	return res

}

func getBase64Char(c int) string {

	if 0 <= c && c <= 25 {

		// 65: A - 90: Z (ASCII)
		return string(rune(c + 65))

	} else if 26 <= c && c <= 51 {

		// 91: a - 122: z (ASCII)
		return string(rune((c - 26) + 97))

	} else if 52 <= c && c <= 61 {

		// converting int to string
		return strconv.Itoa(c - 52)

	} else if c == 62 {

		return "+"

	} else {

		return "/"

	}

}

func HexToBase64(hexStr []byte) string {

	hexBin := []byte{}

	for _, c := range hexStr {
		hexBin = append(hexBin, HexToBin(c)...)
	}

	// fmt.Println(len(hexBin))

	// When the bits cannot be equally divided into 6-bit block, pad the right with 0's
	for len(hexBin)%6 != 0 {

		hexBin = append(hexBin, 0)

	}

	sextets := []int{}

	n := 6
	for i := 0; i < len(hexBin); i += n {
		sixBitBlock := ([6]byte)(hexBin[i : i+n])
		sextets = append(sextets, binToSextet(sixBitBlock))
	}

	// fmt.Println(sextets)

	base64Str := ""

	for _, a := range sextets {
		base64Str = base64Str + getBase64Char(a)
	}

	for len(base64Str)%4 != 0 {

		base64Str = base64Str + "="

	}

	return base64Str
}

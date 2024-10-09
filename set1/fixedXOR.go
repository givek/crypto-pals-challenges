package set1

func xor(x byte, y byte) byte {

	if x == y {

		return 0

	} else {

		return 1

	}

}

func xorByteArr(x []byte, y []byte) []byte {

	// Assuming len(x) == len(y)

	res := []byte{}

	for i := range len(x) {

		res = append(res, xor(x[i], y[i]))

	}

	return res

}

func XorBase16(x []byte, y []byte) string {

	yarr := []byte{}

	for _, h := range y {

		yarr = append(yarr, HexToBin(h)...)

	}

	xarr := []byte{}

	for _, l := range x {

		xarr = append(xarr, HexToBin(l)...)

	}

	xorRes := xorByteArr(xarr, yarr)

	s := ""

	n := 4
	for i := 0; i < len(xorRes); i += n {
		fourBitBlock := ([4]byte)(xorRes[i : i+n])

		v := (fourBitBlock[0] * 8) + (fourBitBlock[1] * 4) + (fourBitBlock[2] * 2) + (fourBitBlock[3] * 1)

		m := map[byte]string{
			0:  "0",
			1:  "1",
			2:  "2",
			3:  "3",
			4:  "4",
			5:  "5",
			6:  "6",
			7:  "7",
			8:  "8",
			9:  "9",
			10: "a",
			11: "b",
			12: "c",
			13: "d",
			14: "e",
			15: "f",
		}

		s = s + m[v]

	}

	return s

}

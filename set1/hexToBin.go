package set1

import (
	"bytes"
)

func HexToBin(c byte) []byte {

	hexToBinMap := map[byte][]byte{
		'0': {0, 0, 0, 0},
		'1': {0, 0, 0, 1},
		'2': {0, 0, 1, 0},
		'3': {0, 0, 1, 1},
		'4': {0, 1, 0, 0},
		'5': {0, 1, 0, 1},
		'6': {0, 1, 1, 0},
		'7': {0, 1, 1, 1},
		'8': {1, 0, 0, 0},
		'9': {1, 0, 0, 1},
		'a': {1, 0, 1, 0},
		'b': {1, 0, 1, 1},
		'c': {1, 1, 0, 0},
		'd': {1, 1, 0, 1},
		'e': {1, 1, 1, 0},
		'f': {1, 1, 1, 1},
	}

	cLower := bytes.ToLower([]byte{c})[0]

	return hexToBinMap[cLower]
}

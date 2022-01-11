package main

import (
	"bytes"
	"encoding/binary"
	"strings"
)

var con = [16]rune{
	'b', 'd', 'f', 'g',
	'h', 'j', 'k', 'l',
	'm', 'n', 'p', 'r',
	's', 't', 'v', 'z',
}

var vow = [4]rune{
	'a', 'i', 'o', 'u',
}

var cd = map[string]uint16{
	"b": 0x0, "d": 0x1, "f": 0x2, "g": 0x3,
	"h": 0x4, "j": 0x5, "k": 0x6, "l": 0x7,
	"m": 0x8, "n": 0x9, "p": 0xA, "r": 0xB,
	"s": 0xC, "t": 0xD, "v": 0xE, "z": 0xF,
}

var vd = map[string]uint16{
	"a": 0x0, "i": 0x1, "o": 0x2, "u": 0x3,
}

var maskLast2 uint16 = 0x3
var maskLast4 uint16 = 0xf
var maskLast8 uint32 = 0xff

// Proquiunt
type Proquint struct{}

//Convert a Bytes.Buffer to Proquint
func (p Proquint) Encode(b bytes.Buffer) string {
	var result []string

	//If Byte length is odd, incrsease by 1 to make it even
	count := b.Len()
	if (count % 2) != 0 {
		count = count + 1
	}

	//Loop over Buffer in chunks of 2 bytes
	for i := 0; i < count/2; i++ {
		var next2 []byte = b.Next(2)

		//Type Cast bytes to uint16
		n := binary.BigEndian.Uint16(next2)

		//Make Proquint
		proquint := []rune{
			con[(n>>0xc)&maskLast4],
			vow[(n>>0xa)&maskLast2],
			con[(n>>0x6)&maskLast4],
			vow[(n>>0x4)&maskLast2],
			con[(n>>0x0)&maskLast4],
		}

		result = append(result, string(proquint))
	}

	return strings.Join(result, "-")
}

//Decodes a Proquint to byte Buffer
func (p Proquint) Decode(s string) *bytes.Buffer {

	var result []byte
	var n uint16
	var buf = make([]byte, 2)

	//Loop over every proquint
	for _, p := range strings.Split(s, "-") {

		//Decode Proquint
		_1 := cd[string(p[0])]
		_2 := vd[string(p[1])]
		_3 := cd[string(p[2])]
		_4 := vd[string(p[3])]
		_5 := cd[string(p[4])]

		n = (_5 << 0) + (_4 << 4) + (_3 << 6) + (_2 << 10) + (_1 << 12)

		//Tye Cast uint16 to bytes
		binary.BigEndian.PutUint16(buf, n)

		//Append to result
		result = append(result, buf...)

		//Reset variables
		n, buf = 0, make([]byte, 2)
	}

	return bytes.NewBuffer(result)
}

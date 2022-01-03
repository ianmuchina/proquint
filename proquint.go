package proquint

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
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

var MASK_LAST2 uint16 = 0x3
var MASK_LAST4 uint16 = 0xf
var MASK_LAST16 uint16 = 0xffff

// Proquiunt
type Proquint struct{}

//Convert a Bytes.Buffer to Proquint
func (p Proquint) encode(b bytes.Buffer) string {
	//Array of Proquints
	var arr []string
	count := b.Len()

	//If Byte length is odd, incrsease by 1 to make it even
	if (count % 2) != 0 {
		count = count + 1
	}

	//Loop over Buffer in chunks of 2 bytes
	for i := 0; i < count/2; i++ {
		var next2 []byte = b.Next(2)

		//Type Cast to Unsigned Integer
		n := binary.BigEndian.Uint16(next2)

		//Generate Proquint
		tmp := []rune{
			con[(n>>12)&MASK_LAST4],
			vow[(n>>10)&MASK_LAST2],
			con[(n>>6)&MASK_LAST4],
			vow[(n>>4)&MASK_LAST2],
			con[(n>>0)&MASK_LAST4],
		}

		arr = append(arr, string(tmp))
	}

	return strings.Join(arr, "-")
}

//Decodes a Proquint to byte Buffer
func (p Proquint) decode(s string) *bytes.Buffer {

	//Array of Proquints
	var arr []string = strings.Split(s, "-")
	//Decoded Data
	var result []byte

	//Placeholder Variables
	var tmp uint16
	var buf = make([]byte, 2)

	for _, p := range arr {

		//Decode Proquint
		_1 := cd[string(p[0])]
		_2 := vd[string(p[1])]
		_3 := cd[string(p[2])]
		_4 := vd[string(p[3])]
		_5 := cd[string(p[4])]
		tmp = (_5 << 0) + (_4 << 4) + (_3 << 6) + (_2 << 10) + (_1 << 12)

		binary.BigEndian.PutUint16(buf, tmp)
		//Append to result
		result = append(result, buf...)
		//Reset variables
		tmp, buf = 0, make([]byte, 2)
	}

	return bytes.NewBuffer(result)
}

type IP struct{ addr net.IP }

//Convert an IPv4 Address to Byte Buffer
func (x IP) asBytes() bytes.Buffer {
	// IPv4 Address as Integer
	var n uint32
	// IPv4 Address as array of octets
	var o [4]uint32

	var buf = make([]byte, 4)

	//Address split to array of strings
	arr := strings.Split(x.addr.String(), ".")

	//Basic Error Checks
	if len(arr) != 4 {
		return bytes.Buffer{}
	}

	for i, n := range arr {

		cur, err := strconv.ParseUint(n, 10, 8)
		//Return Blank if there's an error
		if err != nil {
			return bytes.Buffer{}
		}

		o[i] = uint32(cur)
	}

	n = (o[0] << 24) + (o[1] << 16) + (o[2] << 8) + o[3]

	//Cast Ip to buffer
	binary.BigEndian.PutUint32(buf, n)

	return *bytes.NewBuffer(buf)
}

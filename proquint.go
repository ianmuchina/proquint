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

var MASK_LAST2 uint16 = 0x3
var MASK_LAST4 uint16 = 0xf
var MASK_LAST16 uint16 = 0xffff

// Proquiunt
type Proquint struct{}

type IP struct {
	addr net.IP
}

//Convert a Bytes.Buffer to Proquint
func (p Proquint) Write(b bytes.Buffer) string {
	//Array of Proquints
	var arr []string
	count := b.Len()

	//If Byte length is odd, increase by 1 to make it even
	if (count % 2) != 0 {
		count++
	}

	//Loop over Buffer in chunks of 2 bytes
	for i := 0; i < count/2; i++ {
		var next2 []byte = b.Next(2)

		//Type Cast to Unsigned Integer
		n := binary.BigEndian.Uint16(next2)

		//Generate Proquint
		_1 := (n >> 0) & MASK_LAST4
		_2 := (n >> 4) & MASK_LAST2
		_3 := (n >> 6) & MASK_LAST4
		_4 := (n >> 10) & MASK_LAST2
		_5 := (n >> 12) & MASK_LAST4

		tmp := []rune{con[_5], vow[_4], con[_3], vow[_2], con[_1]}
		arr = append(arr, string(tmp))
	}

	return strings.Join(arr, "-")
}
func (p Proquint) Read(s string) /*bytes.buffer*/ {
	// todo
}

func (x IP) Read() bytes.Buffer {
	//IPv4 Address as Integer
	var n uint32
	//IPv4 Address as array of octets. int32 for easy shifting
	var o [4]uint32

	var buf = make([]byte, 4)

	strArr := strings.Split(x.addr.String(), ".")

	if len(strArr) != 4 {
		return bytes.Buffer{}
	}

	for i, n := range strArr {
		cur, err := strconv.ParseUint(n, 10, 8)
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

//Helper functions to work with ipv4 addresses
package proquint

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type IP struct{}

//Convert IPv4 Address to 4 byte buffer. Currently accepts quad dotted ips
func (x IP) encode(s string) bytes.Buffer {

	ip := net.ParseIP(s)
	// Address as Integer
	var n uint32
	// Address as octet array
	var o [4]uint32

	var result = make([]byte, 4)

	//Split address  to array of strings
	arr := strings.Split(ip.String(), ".")

	//Check if address has more than 4 octets
	if len(arr) != 4 {
		return bytes.Buffer{}
	}

	for i, n := range arr {

		//Numbers are represented as base 10 integers that can only fit in 8 bits
		base, bitsize := 10, 8

		//convert array of strings to array of numbers
		cur, err := strconv.ParseUint(n, base, bitsize)

		//error Parsing string
		if err != nil {
			return bytes.Buffer{}
		}

		o[i] = uint32(cur)
	}

	n = (o[0] << 24) + (o[1] << 16) + (o[2] << 8) + o[3]

	//convert uint32 to 4 bytes
	binary.BigEndian.PutUint32(result, n)

	return *bytes.NewBuffer(result)
}

//Decode byte buffer to IPv4 Address string.
func (x IP) deocde(buf bytes.Buffer) string {

	var n uint32 = binary.BigEndian.Uint32(buf.Bytes())

	var a = []uint32{
		((n >> 0) & maskLast8),
		((n >> 8) & maskLast8),
		((n >> 16) & maskLast8),
		((n >> 24) & maskLast8),
	}

	return fmt.Sprintf("%d.%d.%d.%d", a[3], a[2], a[1], a[0])
}

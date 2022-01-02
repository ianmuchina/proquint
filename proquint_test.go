package proquint

import (
	"bytes"
	"encoding/binary"
	"net"
	"testing"
)

func TestUint32(t *testing.T) {
	var buf = make([]byte, 4)

	var p Proquint
	var n uint32 = 0xbc614e
	want := "bafus-kajav"

	binary.BigEndian.PutUint32(buf, n)
	result := p.Write(*bytes.NewBuffer(buf))

	if result != want {
		t.Errorf("got:%s want: %s", result, want)
	}
}

func TestUint64(t *testing.T) {
	var buf = make([]byte, 8)
	var p Proquint
	var n uint64 = 0xbc614e7f000001

	want := "bafus-kajav-lusab-babad"

	binary.BigEndian.PutUint64(buf, n)
	result := p.Write(*bytes.NewBuffer(buf))

	if result != want {
		t.Errorf("got:%s want: %s", result, want)
	}

}

func TestIP(t *testing.T) {
	var pq Proquint
	var ip IP

	presets := map[string]string{
		"127.0.0.1":      "lusab-babad",
		"63.84.220.193":  "gutih-tugad",
		"63.118.7.35":    "gutuk-bisog",
		"140.98.193.141": "mudof-sakat",
		"64.255.6.200":   "haguz-biram",
		"128.30.52.45":   "mabiv-gibot",
		"147.67.119.2":   "natag-lisaf",
		"212.58.253.68":  "tibup-zujah",
		"216.35.68.215":  "tobog-higil",
		"216.68.232.21":  "todah-vobij",
		"198.81.129.136": "sinid-makam",
		"12.110.110.204": "budov-kuras",
	}

	for key, val := range presets {
		ip.addr = net.ParseIP(key)
		r := pq.Write(ip.Read())

		if r != val {
			t.Errorf("")
		}
	}

}

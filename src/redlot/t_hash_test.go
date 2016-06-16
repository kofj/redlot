package redlot

import (
	"bytes"
	"testing"
)

func TestEncodeHashKey(t *testing.T) {
	name := []byte("name")
	key := []byte("key")
	expect := []byte{0x68, 0x00, 0x00, 0x00, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x00, 0x00, 0x00, 0x03, 0x6b, 0x65, 0x79}
	encoded := encodeHashKey(name, key)
	if !bytes.Equal(expect, encoded) {
		t.Logf("\nexcept: \n\t %v \nencoded: \n\t %v\n", expect, encoded)
		t.Fail()
	}
}

func TestDecodeHashKey(t *testing.T) {
	raw := []byte{0x68, 0x00, 0x00, 0x00, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x00, 0x00, 0x00, 0x03, 0x6b, 0x65, 0x79}
	name, key := decodeHashKey(raw)
	t.Logf("\nexcept: \n\t 0x6e 0x61 0x6d 0x65 \t 0x6b 0x65 0x79 \ndecoded: \n\t % #x \t % #x\n", name, key)
	if !bytes.Equal(name, []byte("name")) || !bytes.Equal(key, []byte("key")) {
		t.Logf("\nexcept: \n\t name \t key \ndecoded: \n\t %v \t %v\n", name, key)
		t.Fail()
	}
}

func TestEncodeHsizeKey(t *testing.T) {
	name := []byte("name")
	expect := []byte{0x48, 0x00, 0x00, 0x00, 0x04, 0x6e, 0x61, 0x6d, 0x65}
	encoded := encodeHsizeKey(name)
	if !bytes.Equal(expect, encoded) {
		t.Logf("\nexcept: \n\t %v \nencoded: \n\t %v\n", expect, encoded)
		t.Fail()
	}
}

func TestDecodeHsizeKey(t *testing.T) {
	raw := []byte{0x48, 0x00, 0x00, 0x00, 0x04, 0x6e, 0x61, 0x6d, 0x65}
	name := decodeHsizeKey(raw)
	if !bytes.Equal([]byte("name"), name) {
		t.Logf("\nexcept: \n\t 0x6e 0x61 0x6d 0x65 \ndecoded: \n\t % #x\n", name)
		t.Fail()
	}
}

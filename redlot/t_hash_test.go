package redlot

import (
	"bytes"
	"testing"
)

func TestEncodeHashKey(t *testing.T) {
	name := []byte("name")
	key := []byte("key")
	expect := []byte{0x68, 0x00, 0x00, 0x00, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x6b, 0x65, 0x79}
	encoded := encodeHashKey(name, key)
	if !bytes.Equal(expect, encoded) {
		t.Logf("\nexcept: \n\t %v \nencoded: \n\t %v\n", expect, encoded)
		t.Fail()
	}
}

func TestDecodeHashKey(t *testing.T) {
	raw := []byte{0x68, 0x00, 0x00, 0x00, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x6b, 0x65, 0x79}
	name, key := decodeHashKey(raw)
	t.Logf("\nexcept: \n\t 0x6e 0x61 0x6d 0x65 \t 0x6b 0x65 0x79 \ndecoded: \n\t % #x \t % #x\n", name, key)
	if !bytes.Equal(name, []byte("name")) || !bytes.Equal(key, []byte("key")) {
		t.Logf("\nexcept: \n\t name \t key \ndecoded: \n\t %v \t %v\n", name, key)
		t.Fail()
	}
}

func TestEncodeHsizeKey(t *testing.T) {
	name := []byte("name")
	expect := []byte{0x48, 0x6e, 0x61, 0x6d, 0x65}
	encoded := encodeHsizeKey(name)
	if !bytes.Equal(expect, encoded) {
		t.Logf("\nexcept: \n\t %v \nencoded: \n\t %v\n", expect, encoded)
		t.Fail()
	}
}

func TestDecodeHsizeKey(t *testing.T) {
	raw := []byte{0x48, 0x6e, 0x61, 0x6d, 0x65}
	name := decodeHsizeKey(raw)
	if !bytes.Equal([]byte("name"), name) {
		t.Logf("\nexcept: \n\t 0x6e 0x61 0x6d 0x65 \ndecoded: \n\t % #x\n", name)
		t.Fail()
	}
}

func TestHashFuncsArgs(t *testing.T) {
	zeroByte := make([][]byte, 0)
	oneByte := make([][]byte, 1)
	twoBytes := make([][]byte, 2)
	threeBytes := make([][]byte, 3)
	// fourByte := make([][]byte, 4)

	// one args methods
	if _, e := Hsize(zeroByte); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hgetall(zeroByte); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hclear(zeroByte); e != errNosArgs {
		t.Fail()
	}

	// two args methods
	if _, e := Hget(oneByte); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hincr(oneByte); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hdel(oneByte); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hexists(oneByte); e != errNosArgs {
		t.Fail()
	}
	if _, e := MultiHget(oneByte); e != errNosArgs {
		t.Fail()
	}
	if _, e := MultiHdel(oneByte); e != errNosArgs {
		t.Fail()
	}

	// theree args methods
	if _, e := Hset(twoBytes); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hlist(twoBytes); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hrlist(twoBytes); e != errNosArgs {
		t.Fail()
	}

	// four args methods
	if _, e := Hkeys(threeBytes); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hscan(threeBytes); e != errNosArgs {
		t.Fail()
	}
	if _, e := Hrscan(threeBytes); e != errNosArgs {
		t.Fail()
	}

}

func TestHashSizeIncr(t *testing.T) {
	name := []byte("hash")

	db.Delete(encodeHsizeKey(name), nil)

	hashSizeIncr(name, 1)
	if b, err := db.Get(encodeHsizeKey(name), nil); bytesToUint32(b) != 1 || err != nil {
		t.Logf("expect hisize is 1, but get: %d\n", bytesToUint32(b))
		t.Fail()
	}

	hashSizeIncr(name, -1)
	if b, err := db.Get(encodeHsizeKey(name), nil); len(b) != 0 || err == nil {
		t.Log("expect hisize is deleted")
		t.Fail()
	}

}

func TestHashFuncs(t *testing.T) {

	Tom := [][]byte{[]byte("Member"), []byte("Tom"), []byte("1001")}
	Amy := [][]byte{[]byte("Member"), []byte("Amy"), []byte("1002")}

	// test hset
	if r, e := Hset(Tom); r != nil || e != nil {
		t.Logf("reply: %v, error: %v\n", r, e)
		t.Fail()
	} else {
		// test hget
		if r, e := Hget(Tom); r.(string) != "1001" || e != nil {
			t.Logf("Hget [Tom] expect 1001, but: %v, error: %v\n", r, e)
			t.Fail()
		}
		// test hsize
		if r, e := Hsize(Tom); r.(int64) != 1 || e != nil {
			t.Logf("Hsize expect 1, but: %d, error: %v\n", r, e)
			t.Fail()
		}
	}

	// test hsize when set same key field.
	Hset(Tom)
	if r, e := Hsize(Tom); r.(int64) != 1 || e != nil {
		t.Logf("Hsize expect 1, but: %d, error: %v\n", r, e)
		t.Fail()
	}

	if r, e := Hset(Amy); r != nil || e != nil {
		t.Logf("reply: %v, error: %v\n", r, e)
		t.Fail()
	} else {
		if r, e := Hget(Amy); r.(string) != "1002" || e != nil {
			t.Logf("Hget [Amy] expect 1002, but: %v, error: %v\n", r, e)
			t.Fail()
		}
		// test hsize when set different field in same hash.
		if r, e := Hsize(Amy); r.(int64) != 2 || e != nil {
			t.Logf("Hsize expect 1, but: %d, error: %v\n", r, e)
			t.Fail()
		}
	}

	// test hdel and hexists
	if r, e := Hexists(Amy); r.(int64) != 1 || e != nil {
		t.Logf("Hexists [Amy] expect 1, but: %d, error: %v\n", r, e)
		t.Fail()
	}
	Hdel(Tom)
	if r, e := Hexists(Tom); r.(int64) != 0 || e != nil {
		t.Logf("Hexists [Tom] expect 0, but: %d, error: %v\n", r, e)
		t.Fail()
	}
	Hdel(Amy)
	if r, e := Hsize(Amy); r.(int64) != -1 || e != nil {
		t.Logf("Hsize expect -1, but: %d, error: %v\n", r, e)
	}

}

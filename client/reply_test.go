package client

import (
	"bytes"
	"math"
	"testing"
)

func TestBytes(t *testing.T) {
	r := client.Cmd("get", "k").Bytes()

	if !bytes.Equal(r, []byte("v")) {
		t.Logf("expect bytes [% #x], but get[ % #x]\n", []byte("v"), r)
		t.Fail()
	}
}

func TestString(t *testing.T) {
	r := client.Cmd("get", "k").String()

	if r != "v" {
		t.Logf("expect string [v], but get[ %s]\n", r)
		t.Fail()
	}

}

func TestInt(t *testing.T) {
	r := client.Cmd("get", "k").Int()

	if r != 0 {
		t.Logf("expect int [0], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int", "9223372036854775807")
	r = client.Cmd("get", "int").Int()

	if r != 9223372036854775807 {
		t.Logf("expect int [9223372036854775807], but get[ %d]\n", r)
		t.Fail()
	}

	// test value out of range
	client.Cmd("set", "int", "9223372036854775808")
	r = client.Cmd("get", "int").Int()

	if r != 9223372036854775807 {
		t.Logf("expect int [9223372036854775807], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestInt8(t *testing.T) {
	client.Cmd("set", "int8", "127")
	r := client.Cmd("get", "int8").Int8()

	if r != 127 {
		t.Logf("expect int8 [127], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int8", "-128")
	r = client.Cmd("get", "int8").Int8()

	if r != -128 {
		t.Logf("expect int8 [-128], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int8", "128")
	r = client.Cmd("get", "int8").Int8()

	if r != 127 {
		t.Logf("expect int8 [127], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestInt16(t *testing.T) {
	client.Cmd("set", "int16", "32767")
	r := client.Cmd("get", "int16").Int16()

	if r != 32767 {
		t.Logf("expect int16 [32767], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int16", "-32768")
	r = client.Cmd("get", "int16").Int16()

	if r != -32768 {
		t.Logf("expect int16 [-32768], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int16", "32768")
	r = client.Cmd("get", "int16").Int16()

	if r != 32767 {
		t.Logf("expect int16 [32767], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestInt32(t *testing.T) {
	client.Cmd("set", "int32", "2147483647")
	r := client.Cmd("get", "int32").Int32()

	if r != 2147483647 {
		t.Logf("expect int32 [2147483647], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int32", "-2147483648")
	r = client.Cmd("get", "int32").Int32()

	if r != -2147483648 {
		t.Logf("expect int32 [-2147483648], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int32", "2147483648")
	r = client.Cmd("get", "int32").Int32()

	if r != 2147483647 {
		t.Logf("expect int32 [2147483647], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestInt64(t *testing.T) {
	client.Cmd("set", "int64", "9223372036854775807")
	r := client.Cmd("get", "int64").Int64()

	if r != 9223372036854775807 {
		t.Logf("expect int64 [9223372036854775807], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int64", "-9223372036854775808")
	r = client.Cmd("get", "int64").Int64()

	if r != -9223372036854775808 {
		t.Logf("expect int64 [-9223372036854775808], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int64", "9223372036854775808")
	r = client.Cmd("get", "int64").Int64()

	if r != 9223372036854775807 {
		t.Logf("expect int64 [9223372036854775807], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestUint8(t *testing.T) {
	client.Cmd("set", "uint8", "255")
	r := client.Cmd("get", "int64").Uint8()

	if r != 255 {
		t.Logf("expect uint8 [255], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint8", "256")
	r = client.Cmd("get", "uint8").Uint8()

	if r != 255 {
		t.Logf("expect uint8 [255], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint8", "-1")
	r = client.Cmd("get", "uint8").Uint8()

	if r != 0 {
		t.Logf("expect uint8 [0], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestUint16(t *testing.T) {
	client.Cmd("set", "uint16", "65535")
	r := client.Cmd("get", "int64").Uint16()

	if r != 65535 {
		t.Logf("expect uint16 [65535], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint16", "65536")
	r = client.Cmd("get", "uint16").Uint16()

	if r != 65535 {
		t.Logf("expect uint16 [65535], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint16", "-1")
	r = client.Cmd("get", "uint16").Uint16()

	if r != 0 {
		t.Logf("expect uint16 [0], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestUint32(t *testing.T) {
	client.Cmd("set", "uint32", "4294967295")
	r := client.Cmd("get", "int64").Uint32()

	if r != 4294967295 {
		t.Logf("expect uint32 [4294967295], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint32", "4294967296")
	r = client.Cmd("get", "uint32").Uint32()

	if r != 4294967295 {
		t.Logf("expect uint32 [4294967295], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint32", "-1")
	r = client.Cmd("get", "uint32").Uint32()

	if r != 0 {
		t.Logf("expect uint32 [0], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestUint64(t *testing.T) {
	client.Cmd("set", "uint64", "18446744073709551615")
	r := client.Cmd("get", "uint64").Uint64()

	if r != 18446744073709551615 {
		t.Logf("expect uint64 [18446744073709551615], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint64", "18446744073709551616")
	r = client.Cmd("get", "uint64").Uint64()

	if r != 18446744073709551615 {
		t.Logf("expect uint64 [18446744073709551615], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint64", "-1")
	r = client.Cmd("get", "uint64").Uint64()

	if r != 0 {
		t.Logf("expect uint64 [0], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestUint(t *testing.T) {
	client.Cmd("set", "uint", "18446744073709551615")
	r := client.Cmd("get", "uint").Uint()

	if r != 18446744073709551615 {
		t.Logf("expect uint [18446744073709551615], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint", "18446744073709551616")
	r = client.Cmd("get", "uint").Uint()

	if r != 18446744073709551615 {
		t.Logf("expect uint [18446744073709551615], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "uint", "-1")
	r = client.Cmd("get", "uint").Uint()

	if r != 0 {
		t.Logf("expect uint [0], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestBool(t *testing.T) {
	client.Cmd("set", "bool", "true")
	r := client.Cmd("get", "bool").Bool()

	if r != true {
		t.Logf("expect bool [true], but get %b\n", r)
		t.Fail()
	}

	client.Cmd("set", "bool", "1")
	r = client.Cmd("get", "bool").Bool()

	if r != true {
		t.Logf("expect bool [true], but get %t\n", r)
		t.Fail()
	}

}

func TestFloat32(t *testing.T) {
	client.Cmd("set", "float32", "7.22")
	r := client.Cmd("get", "float32").Float32()

	if r != 7.22 {
		t.Logf("expect float32 [7.22], but get %b\n", r)
		t.Fail()
	}

	client.Cmd("set", "float32", "inf")
	r = client.Cmd("get", "float32").Float32()

	if !math.IsInf(float64(r), 1) {
		t.Logf("expect float32 [+Inf], but get %b\n", r)
		t.Fail()
	}

	client.Cmd("set", "float32", "NaN")
	r = client.Cmd("get", "float32").Float32()

	if !math.IsNaN(float64(r)) {
		t.Logf("expect float32 [NaN], but get %b\n", r)
		t.Fail()
	}

}

func TestFloat64(t *testing.T) {
	client.Cmd("set", "float64", "15.95")
	r := client.Cmd("get", "float64").Float64()

	if r != 15.95 {
		t.Logf("expect float64 [15.95], but get %b\n", r)
		t.Fail()
	}

	client.Cmd("set", "float64", "inf")
	r = client.Cmd("get", "float64").Float64()

	if !math.IsInf(r, 1) {
		t.Logf("expect float64 [+Inf], but get %b\n", r)
		t.Fail()
	}

	client.Cmd("set", "float64", "NaN")
	r = client.Cmd("get", "float64").Float64()

	if !math.IsNaN(r) {
		t.Logf("expect float64 [NaN], but get %b\n", r)
		t.Fail()
	}

}

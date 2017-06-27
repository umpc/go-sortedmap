package asc

import "testing"

func TestUint8(t *testing.T) {
	if Uint8(uint8(1), uint8(0)) {
		t.Fatal("asc.TestUint8 failed: %v", greaterThanErr)
	}
}

func TestUint16(t *testing.T) {
	if Uint16(uint16(1), uint16(0)) {
		t.Fatal("asc.TestUint16 failed: %v", greaterThanErr)
	}
}

func TestUint32(t *testing.T) {
	if Uint32(uint32(1), uint32(0)) {
		t.Fatal("asc.TestUint32 failed: %v", greaterThanErr)
	}
}

func TestUint64(t *testing.T) {
	if Uint64(uint64(1), uint64(0)) {
		t.Fatal("asc.TestUint64 failed: %v", greaterThanErr)
	}
}

func TestInt8(t *testing.T) {
	if Int8(int8(1), int8(0)) {
		t.Fatal("asc.TestInt8 failed: %v", greaterThanErr)
	}
}

func TestInt16(t *testing.T) {
	if Int16(int16(1), int16(0)) {
		t.Fatal("asc.TestInt16 failed: %v", greaterThanErr)
	}
}

func TestInt32(t *testing.T) {
	if Int32(int32(1), int32(0)) {
		t.Fatal("asc.TestInt32 failed: %v", greaterThanErr)
	}
}

func TestInt64(t *testing.T) {
	if Int64(int64(1), int64(0)) {
		t.Fatal("asc.TestInt64 failed: %v", greaterThanErr)
	}
}

func TestFloat32(t *testing.T) {
	if Float32(float32(1), float32(0)) {
		t.Fatal("asc.TestFloat32 failed: %v", greaterThanErr)
	}
}

func TestFloat64(t *testing.T) {
	if Float64(float64(1), float64(0)) {
		t.Fatal("asc.TestFloat64 failed: %v", greaterThanErr)
	}
}

func TestUint(t *testing.T) {
	if Uint(uint(1), uint(0)) {
		t.Fatal("asc.TestUint failed: %v", greaterThanErr)
	}
}

func TestInt(t *testing.T) {
	if Int(int(1), 0) {
		t.Fatal("asc.TestInt failed: %v", greaterThanErr)
	}
}
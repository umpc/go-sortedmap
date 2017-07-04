package desc

import "testing"

func TestUint8(t *testing.T) {
	if Uint8(uint8(0), uint8(1)) {
		t.Fatalf("desc.TestUint8 failed: %v\n", greaterThanErr)
	}
}

func TestUint16(t *testing.T) {
	if Uint16(uint16(0), uint16(1)) {
		t.Fatalf("desc.TestUint16 failed: %v\n", greaterThanErr)
	}
}

func TestUint32(t *testing.T) {
	if Uint32(uint32(0), uint32(1)) {
		t.Fatalf("desc.TestUint32 failed: %v\n", greaterThanErr)
	}
}

func TestUint64(t *testing.T) {
	if Uint64(uint64(0), uint64(1)) {
		t.Fatalf("desc.TestUint64 failed: %v\n", greaterThanErr)
	}
}

func TestInt8(t *testing.T) {
	if Int8(int8(0), int8(1)) {
		t.Fatalf("desc.TestInt8 failed: %v\n", greaterThanErr)
	}
}

func TestInt16(t *testing.T) {
	if Int16(int16(0), int16(1)) {
		t.Fatalf("desc.TestInt16 failed: %v\n", greaterThanErr)
	}
}

func TestInt32(t *testing.T) {
	if Int32(int32(0), int32(1)) {
		t.Fatalf("desc.TestInt32 failed: %v\n", greaterThanErr)
	}
}

func TestInt64(t *testing.T) {
	if Int64(int64(0), int64(1)) {
		t.Fatalf("desc.TestInt64 failed: %v\n", greaterThanErr)
	}
}

func TestFloat32(t *testing.T) {
	if Float32(float32(0), float32(1)) {
		t.Fatalf("desc.TestFloat32 failed: %v\n", greaterThanErr)
	}
}

func TestFloat64(t *testing.T) {
	if Float64(float64(0), float64(1)) {
		t.Fatalf("desc.TestFloat64 failed: %v\n", greaterThanErr)
	}
}

func TestUint(t *testing.T) {
	if Uint(uint(0), uint(1)) {
		t.Fatalf("desc.TestUint failed: %v\n", greaterThanErr)
	}
}

func TestInt(t *testing.T) {
	if Int(int(0), 0) {
		t.Fatalf("desc.TestInt failed: %v\n", greaterThanErr)
	}
}

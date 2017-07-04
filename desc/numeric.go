package desc

// Uint8 is a greater than comparison function for the Uint8 numeric type.
func Uint8(i, j interface{}) bool {
	return i.(uint8) > j.(uint8)
}

// Uint16 is a greater than comparison function for the Uint16 numeric type.
func Uint16(i, j interface{}) bool {
	return i.(uint16) > j.(uint16)
}

// Uint32 is a greater than comparison function for the Uint32 numeric type.
func Uint32(i, j interface{}) bool {
	return i.(uint32) > j.(uint32)
}

// Uint64 is a greater than comparison function for the Uint64 numeric type.
func Uint64(i, j interface{}) bool {
	return i.(uint64) > j.(uint64)
}

// Int8 is a greater than comparison function for the Int8 numeric type.
func Int8(i, j interface{}) bool {
	return i.(int8) > j.(int8)
}

// Int16 is a greater than comparison function for the Int16 numeric type.
func Int16(i, j interface{}) bool {
	return i.(int16) > j.(int16)
}

// Int32 is a greater than comparison function for the Int32 numeric type.
func Int32(i, j interface{}) bool {
	return i.(int32) > j.(int32)
}

// Int64 is a greater than comparison function for the Int64 numeric type.
func Int64(i, j interface{}) bool {
	return i.(int64) > j.(int64)
}

// Float32 is a greater than comparison function for the Float32 numeric type.
func Float32(i, j interface{}) bool {
	return i.(float32) > j.(float32)
}

// Float64 is a greater than comparison function for the Float64 numeric type.
func Float64(i, j interface{}) bool {
	return i.(float64) > j.(float64)
}

// Uint is a greater than comparison function for the Uint numeric type.
func Uint(i, j interface{}) bool {
	return i.(uint) > j.(uint)
}

// Int is a greater than comparison function for the Int numeric type.
func Int(i, j interface{}) bool {
	return i.(int) > j.(int)
}

package asc

func Uint8(i, j interface{}) bool {
	return i.(uint8) < j.(uint8)
}

func Uint16(i, j interface{}) bool {
	return i.(uint16) < j.(uint16)
}

func Uint32(i, j interface{}) bool {
	return i.(uint32) < j.(uint32)
}

func Uint64(i, j interface{}) bool {
	return i.(uint64) < j.(uint64)
}

func Int8(i, j interface{}) bool {
	return i.(int8) < j.(int8)
}

func Int16(i, j interface{}) bool {
	return i.(int16) < j.(int16)
}

func Int32(i, j interface{}) bool {
	return i.(int32) < j.(int32)
}

func Int64(i, j interface{}) bool {
	return i.(int64) < j.(int64)
}

func Float32(i, j interface{}) bool {
	return i.(float32) < j.(float32)
}

func Float64(i, j interface{}) bool {
	return i.(float64) < j.(float64)
}

func Uint(i, j interface{}) bool {
	return i.(uint) < j.(uint)
}

func Int(i, j interface{}) bool {
	return i.(int) < j.(int)
}
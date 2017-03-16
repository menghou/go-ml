package base

import (
	"math"
	"unsafe"
)

func PackFloatToBytes(val float64) []byte {
	u64 := math.Float64bits(val)
	re := make([]byte, 8)
	re[7] = byte(u64 & (0xFF << 56) >> 56)
	re[6] = byte(u64 & (0xFF << 48) >> 48)
	re[5] = byte(u64 & (0xFF << 40) >> 40)
	re[4] = byte(u64 & (0xFF << 32) >> 32)
	re[3] = byte(u64 & (0xFF << 24) >> 24)
	re[2] = byte(u64 & (0xFF << 16) >> 16)
	re[1] = byte(u64 & (0xFF << 8) >> 8)
	re[0] = byte(u64 & (0xFF << 0) >> 0)
	return re
}

func UnPackBytesToFloat(rawVal []byte) float64 {
	return *(*float64)(unsafe.Pointer(&rawVal[0]))
}

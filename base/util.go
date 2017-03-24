package base

import (
	"math"
	"unsafe"
)

func PackFloatToBytes(val float64) []byte {
	return PackU64ToBytes(math.Float64bits(val))
}
func PackU64ToBytes(val uint64) []byte {
	re := make([]byte, 8)
	re[7] = byte(val & (0xFF << 56) >> 56)
	re[6] = byte(val & (0xFF << 48) >> 48)
	re[5] = byte(val & (0xFF << 40) >> 40)
	re[4] = byte(val & (0xFF << 32) >> 32)
	re[3] = byte(val & (0xFF << 24) >> 24)
	re[2] = byte(val & (0xFF << 16) >> 16)
	re[1] = byte(val & (0xFF << 8) >> 8)
	re[0] = byte(val & (0xFF << 0) >> 0)
	return re
}
func UnPackBytesToFloat(rawVal []byte) float64 {
	return *(*float64)(unsafe.Pointer(&rawVal[0]))
}

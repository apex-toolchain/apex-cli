package util

type BitSize int

const (
	Bits8  BitSize = 8
	Bits16         = 16
	Bits32         = 32
	Bits64         = 64
)

func PackUintLE(n uint64, bitSize BitSize) []byte {
	if bitSize <= 0 || bitSize > 64 {
		panic("bitSize out of range")
	}

	mask := uint64(1<<bitSize) - 1
	trimmed := n & mask

	byteLen := int((bitSize + 7) / 8)
	out := make([]byte, byteLen)

	// Little-endian: lowest byte first
	for i := 0; i < byteLen; i++ {
		out[i] = byte(trimmed & 0xFF)
		trimmed >>= 8
	}

	return out
}

func UnpackUintLE(data []byte, bitSize BitSize) uint64 {
	if bitSize <= 0 || bitSize > 64 {
		panic("bitSize out of range")
	}

	byteLen := int((bitSize + 7) / 8)
	if len(data) != byteLen {
		panic("invalid byte slice length for given bitSize")
	}

	var n uint64 = 0
	for i := byteLen - 1; i >= 0; i-- {
		n <<= 8
		n |= uint64(data[i])
	}

	if bitSize < 64 {
		n &= (uint64(1)<<bitSize - 1)
	}

	return n
}

func PackIntLE(n int64, bitSize BitSize) []byte {
	if bitSize <= 0 || bitSize > 64 {
		panic("bitSize out of range")
	}

	mask := uint64(1<<bitSize) - 1
	trimmed := uint64(n) & mask

	byteLen := int((bitSize + 7) / 8)
	out := make([]byte, byteLen)

	for i := 0; i < byteLen; i++ {
		out[i] = byte(trimmed & 0xFF)
		trimmed >>= 8
	}

	return out
}

func UnpackIntLE(data []byte, bitSize BitSize) int64 {
	if bitSize <= 0 || bitSize > 64 {
		panic("bitSize out of range")
	}

	byteLen := int((bitSize + 7) / 8)
	if len(data) != byteLen {
		panic("invalid byte slice length for given bitSize")
	}

	var n uint64 = 0
	for i := byteLen - 1; i >= 0; i-- {
		n <<= 8
		n |= uint64(data[i])
	}

	if bitSize < 64 {
		signBit := uint64(1) << (bitSize - 1)
		mask := (uint64(1) << bitSize) - 1
		n &= mask
		if n&signBit != 0 {
			n |= ^mask
		}
	}

	return int64(n)
}

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

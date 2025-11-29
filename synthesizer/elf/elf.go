package elf

type BitSize int

const (
	Bit32 BitSize = 32
	Bit64 BitSize = 64
)

type ElfWriter struct {
	BitSize BitSize
}

func NewElfWriter(bitSize BitSize) *ElfWriter {
	return &ElfWriter{
		BitSize: bitSize,
	}
}

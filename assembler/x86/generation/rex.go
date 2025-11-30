package generation

type Rex struct {
	Needed bool // whether to emit a REX byte at all
	W      byte
	R      byte
	X      byte
	B      byte
}

func MakeRexByte(rex Rex) byte {
	return (0x4 << 4) | ((rex.W & 1) << 3) | ((rex.R & 1) << 2) | ((rex.X & 1) << 1) | (rex.B & 1)
}

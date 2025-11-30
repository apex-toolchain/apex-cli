package generation

func make_rexbyte(w, r, x, b byte) byte {
	//generates a REX / Register Extension prefix byte, wrxb must be 0b0 or 0b1
	return (4 << 4) | (w << 3) | (r << 2) | (x << 1) | b
}

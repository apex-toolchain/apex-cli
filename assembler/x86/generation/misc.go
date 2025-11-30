package generation

// Miscellaneous opcode registrations (NOP, etc.).
func init() {
	opcodeTable["NOP"] = OpcodeMeta{
		Bytes:       []byte{0x90},
		Operands:    []OpcodeType{},
		Order:       []BinaryOrder{BaseOpcode},
		Flags:       FlagDirect,
		Description: "Does nothing.",
	}
}

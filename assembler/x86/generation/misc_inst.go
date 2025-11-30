package generation

// Miscellaneous opcode registrations (moved back into generation package).
func init() {
	opcodeTable["NOP"] = OpcodeMeta{
		Bytes:       []byte{0x90},
		Operands:    []OpcodeType{},
		Order:       []BinaryOrder{BaseOpcode},
		Flags:       FlagDirect,
		Description: "Does nothing.",
	}
}

package generation

// XOR-related opcode registrations.
func init() {
	opcodeTable["XOR_r64_r64"] = OpcodeMeta{
		Bytes:       []byte{0x33},
		Operands:    []OpcodeType{OP_R64, OP_R64},
		Order:       []BinaryOrder{BaseOpcode, SetMod, RegOpFromOp0, SrcInRM},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "XOR r64 with r64 using ModRM (register-to-register)",
	}
}

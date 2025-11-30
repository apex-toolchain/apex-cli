package generation

// ADD-related opcode registrations (moved back into generation package).
func init() {
	opcodeTable["ADD_r64_imm32"] = OpcodeMeta{
		Bytes:       []byte{0x81},
		Operands:    []OpcodeType{OP_R64, OP_Imm32},
		Order:       []BinaryOrder{BaseOpcode, SetMod, Soe0, DestInRM, StoreImmediate},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Add a 32 bit immediate to a 64 bit register.",
	}
}

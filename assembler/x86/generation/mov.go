package generation

// MOV-related opcode registrations. These populate the shared opcodeTable at
// package init time.
func init() {
	opcodeTable["MOV_r64_imm64"] = OpcodeMeta{
		Bytes:       []byte{0xB8},
		Operands:    []OpcodeType{OP_R64, OP_Imm64},
		Order:       []BinaryOrder{BaseOpcode, AddLow3, InstantUImm},
		Flags:       FlagRexW | FlagRex | FlagDirect,
		Description: "Move a 64 bit immediate into a 64 bit register. (REX.W + 8B /r)",
	}

	opcodeTable["MOV_r64_rm64"] = OpcodeMeta{
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R64, OP_RM64},
		Order:       []BinaryOrder{BaseOpcode, RegOpFromOp0, SrcInRM},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Move r/m64 into r64 (0x8B /r)",
	}

	opcodeTable["MOV_r64_r64"] = OpcodeMeta{
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R64, OP_R64},
		Order:       []BinaryOrder{BaseOpcode, SetMod, RegOpFromOp0, SrcInRM},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Move r64 into r64 using ModRM (register to register) (0x8B /r)",
	}

	opcodeTable["MOV_rm64_r64"] = OpcodeMeta{
		Bytes:       []byte{0x89},
		Operands:    []OpcodeType{OP_RM64, OP_R64},
		Order:       []BinaryOrder{BaseOpcode, SetMod, RegOpFromOp1, DestInRM},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Move r64 into r/m64 (0x89 /r)",
	}

	opcodeTable["MOV_r32_imm32"] = OpcodeMeta{
		Bytes:       []byte{0xB8},
		Operands:    []OpcodeType{OP_R32, OP_Imm32},
		Order:       []BinaryOrder{BaseOpcode, AddLow3, StoreImmediate},
		Flags:       FlagDirect,
		Description: "Move a 32 bit immediate into a 32 bit register (B8+rd)",
	}

	opcodeTable["MOV_r32_rm32"] = OpcodeMeta{
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R32, OP_RM32},
		Order:       []BinaryOrder{BaseOpcode, RegOpFromOp0, SrcInRM},
		Flags:       FlagModRM,
		Description: "Move r/m32 into r32 (0x8B /r)",
	}

	opcodeTable["MOV_r32_r32"] = OpcodeMeta{
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R32, OP_R32},
		Order:       []BinaryOrder{BaseOpcode, SetMod, RegOpFromOp0, SrcInRM},
		Flags:       FlagModRM,
		Description: "Move r32 into r32 using ModRM (register to register) (0x8B /r)",
	}

	opcodeTable["MOV_rm32_r32"] = OpcodeMeta{
		Bytes:       []byte{0x89},
		Operands:    []OpcodeType{OP_RM32, OP_R32},
		Order:       []BinaryOrder{BaseOpcode, SetMod, RegOpFromOp1, DestInRM},
		Flags:       FlagModRM,
		Description: "Move r32 into r/m32 (0x89 /r)",
	}

	opcodeTable["MOV_r8_rm8"] = OpcodeMeta{
		Bytes:       []byte{0x8A},
		Operands:    []OpcodeType{OP_R8, OP_RM8},
		Order:       []BinaryOrder{BaseOpcode, RegOpFromOp0, SrcInRM},
		Flags:       FlagModRM,
		Description: "Move r/m8 into r8 (0x8A /r)",
	}

	opcodeTable["MOV_rm8_r8"] = OpcodeMeta{
		Bytes:       []byte{0x88},
		Operands:    []OpcodeType{OP_RM8, OP_R8},
		Order:       []BinaryOrder{BaseOpcode, SetMod, RegOpFromOp1, DestInRM},
		Flags:       FlagModRM,
		Description: "Move r8 into r/m8 (0x88 /r)",
	}

	opcodeTable["MOV_rm32_imm32"] = OpcodeMeta{
		Bytes:       []byte{0xC7},
		Operands:    []OpcodeType{OP_RM32, OP_Imm32},
		Order:       []BinaryOrder{BaseOpcode, Soe0, DestInRM, StoreImmediate},
		Flags:       FlagModRM,
		Description: "Move a 32 bit immediate into r/m32 (C7 /0)",
	}

	opcodeTable["MOV_rm8_imm8"] = OpcodeMeta{
		Bytes:       []byte{0xC6},
		Operands:    []OpcodeType{OP_RM8, OP_Imm8},
		Order:       []BinaryOrder{BaseOpcode, Soe0, DestInRM, InstantUImm},
		Flags:       FlagModRM,
		Description: "Move an 8 bit immediate into r/m8 (C6 /0)",
	}

	opcodeTable["MOV_r8_imm8"] = OpcodeMeta{
		Bytes:       []byte{0xB0},
		Operands:    []OpcodeType{OP_R8, OP_Imm8},
		Order:       []BinaryOrder{BaseOpcode, AddLow3, InstantUImm},
		Flags:       FlagDirect,
		Description: "Move an 8 bit immediate into an 8 bit register (B0+rd)",
	}

	opcodeTable["MOV_rm64_imm32"] = OpcodeMeta{
		Bytes:       []byte{0xC7},
		Operands:    []OpcodeType{OP_RM64, OP_Imm32},
		Order:       []BinaryOrder{BaseOpcode, Soe0, DestInRM, StoreImmediate},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Move a 32 bit immediate into r/m64 (C7 /0)",
	}
}

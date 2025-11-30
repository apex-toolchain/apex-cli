package instructions

import "chasm/assembler/x86/generation"

func init() {
	generation.RegisterOpcode("MOV_r64_imm64", generation.OpcodeMeta{
		Bytes:       []byte{0xB8},
		Operands:    []generation.OpcodeType{generation.OP_R64, generation.OP_Imm64},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.AddLow3, generation.InstantUImm},
		Flags:       generation.FlagRexW | generation.FlagRex | generation.FlagDirect,
		Description: "Move a 64 bit immediate into a 64 bit register. (REX.W + 8B /r)",
	})

	generation.RegisterOpcode("MOV_r64_rm64", generation.OpcodeMeta{
		Bytes:       []byte{0x8B},
		Operands:    []generation.OpcodeType{generation.OP_R64, generation.OP_RM64},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.RegOpFromOp0, generation.SrcInRM},
		Flags:       generation.FlagRexW | generation.FlagRex | generation.FlagModRM,
		Description: "Move r/m64 into r64 (0x8B /r)",
	})

	generation.RegisterOpcode("MOV_r64_r64", generation.OpcodeMeta{
		Bytes:       []byte{0x8B},
		Operands:    []generation.OpcodeType{generation.OP_R64, generation.OP_R64},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.SetMod, generation.RegOpFromOp0, generation.SrcInRM},
		Flags:       generation.FlagRexW | generation.FlagRex | generation.FlagModRM,
		Description: "Move r64 into r64 using ModRM (register to register) (0x8B /r)",
	})

	generation.RegisterOpcode("MOV_rm64_r64", generation.OpcodeMeta{
		Bytes:       []byte{0x89},
		Operands:    []generation.OpcodeType{generation.OP_RM64, generation.OP_R64},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.SetMod, generation.RegOpFromOp1, generation.DestInRM},
		Flags:       generation.FlagRexW | generation.FlagRex | generation.FlagModRM,
		Description: "Move r64 into r/m64 (0x89 /r)",
	})

	generation.RegisterOpcode("MOV_r32_imm32", generation.OpcodeMeta{
		Bytes:       []byte{0xB8},
		Operands:    []generation.OpcodeType{generation.OP_R32, generation.OP_Imm32},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.AddLow3, generation.StoreImmediate},
		Flags:       generation.FlagDirect,
		Description: "Move a 32 bit immediate into a 32 bit register (B8+rd)",
	})

	generation.RegisterOpcode("MOV_r32_rm32", generation.OpcodeMeta{
		Bytes:       []byte{0x8B},
		Operands:    []generation.OpcodeType{generation.OP_R32, generation.OP_RM32},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.RegOpFromOp0, generation.SrcInRM},
		Flags:       generation.FlagModRM,
		Description: "Move r/m32 into r32 (0x8B /r)",
	})

	generation.RegisterOpcode("MOV_r32_r32", generation.OpcodeMeta{
		Bytes:       []byte{0x8B},
		Operands:    []generation.OpcodeType{generation.OP_R32, generation.OP_R32},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.SetMod, generation.RegOpFromOp0, generation.SrcInRM},
		Flags:       generation.FlagModRM,
		Description: "Move r32 into r32 using ModRM (register to register) (0x8B /r)",
	})

	generation.RegisterOpcode("MOV_rm32_r32", generation.OpcodeMeta{
		Bytes:       []byte{0x89},
		Operands:    []generation.OpcodeType{generation.OP_RM32, generation.OP_R32},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.SetMod, generation.RegOpFromOp1, generation.DestInRM},
		Flags:       generation.FlagModRM,
		Description: "Move r32 into r/m32 (0x89 /r)",
	})

	generation.RegisterOpcode("MOV_r8_rm8", generation.OpcodeMeta{
		Bytes:       []byte{0x8A},
		Operands:    []generation.OpcodeType{generation.OP_R8, generation.OP_RM8},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.RegOpFromOp0, generation.SrcInRM},
		Flags:       generation.FlagModRM,
		Description: "Move r/m8 into r8 (0x8A /r)",
	})

	generation.RegisterOpcode("MOV_rm8_r8", generation.OpcodeMeta{
		Bytes:       []byte{0x88},
		Operands:    []generation.OpcodeType{generation.OP_RM8, generation.OP_R8},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.SetMod, generation.RegOpFromOp1, generation.DestInRM},
		Flags:       generation.FlagModRM,
		Description: "Move r8 into r/m8 (0x88 /r)",
	})

	generation.RegisterOpcode("MOV_rm32_imm32", generation.OpcodeMeta{
		Bytes:       []byte{0xC7},
		Operands:    []generation.OpcodeType{generation.OP_RM32, generation.OP_Imm32},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.Soe0, generation.DestInRM, generation.StoreImmediate},
		Flags:       generation.FlagModRM,
		Description: "Move a 32 bit immediate into r/m32 (C7 /0)",
	})

	generation.RegisterOpcode("MOV_rm8_imm8", generation.OpcodeMeta{
		Bytes:       []byte{0xC6},
		Operands:    []generation.OpcodeType{generation.OP_RM8, generation.OP_Imm8},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.Soe0, generation.DestInRM, generation.InstantUImm},
		Flags:       generation.FlagModRM,
		Description: "Move an 8 bit immediate into r/m8 (C6 /0)",
	})

	generation.RegisterOpcode("MOV_r8_imm8", generation.OpcodeMeta{
		Bytes:       []byte{0xB0},
		Operands:    []generation.OpcodeType{generation.OP_R8, generation.OP_Imm8},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.AddLow3, generation.InstantUImm},
		Flags:       generation.FlagDirect,
		Description: "Move an 8 bit immediate into an 8 bit register (B0+rd)",
	})

	generation.RegisterOpcode("MOV_rm64_imm32", generation.OpcodeMeta{
		Bytes:       []byte{0xC7},
		Operands:    []generation.OpcodeType{generation.OP_RM64, generation.OP_Imm32},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.Soe0, generation.DestInRM, generation.StoreImmediate},
		Flags:       generation.FlagRexW | generation.FlagRex | generation.FlagModRM,
		Description: "Move a 32 bit immediate into r/m64 (C7 /0)",
	})
}

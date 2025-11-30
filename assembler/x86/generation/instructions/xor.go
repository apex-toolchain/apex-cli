package instructions

import "chasm/assembler/x86/generation"

// XOR-related opcode registrations.
func init() {
	generation.RegisterOpcode("XOR_r64_r64", generation.OpcodeMeta{
		Bytes:       []byte{0x33},
		Operands:    []generation.OpcodeType{generation.OP_R64, generation.OP_R64},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.SetMod, generation.RegOpFromOp0, generation.SrcInRM},
		Flags:       generation.FlagRexW | generation.FlagRex | generation.FlagModRM,
		Description: "XOR r64 with r64 using ModRM (register-to-register)",
	})
}

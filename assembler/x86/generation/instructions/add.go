package instructions

import "chasm/assembler/x86/generation"

// ADD-related opcode registrations.
func init() {
	generation.RegisterOpcode("ADD_r64_imm32", generation.OpcodeMeta{
		Bytes:       []byte{0x81},
		Operands:    []generation.OpcodeType{generation.OP_R64, generation.OP_Imm32},
		Order:       []generation.BinaryOrder{generation.BaseOpcode, generation.SetMod, generation.Soe0, generation.DestInRM, generation.StoreImmediate},
		Flags:       generation.FlagRexW | generation.FlagRex | generation.FlagModRM,
		Description: "Add a 32 bit immediate to a 64 bit register.",
	})
}

package instructions

import "chasm/assembler/x86/generation"

// Miscellaneous opcode registrations (NOP, etc.).
func init() {
	generation.RegisterOpcode("NOP", generation.OpcodeMeta{
		Bytes:       []byte{0x90},
		Operands:    []generation.OpcodeType{},
		Order:       []generation.BinaryOrder{generation.BaseOpcode},
		Flags:       generation.FlagDirect,
		Description: "Does nothing.",
	})
}

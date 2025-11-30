package generation

import (
	"chasm/util"
)

type OpcodeMeta struct {
	Bytes       []byte       // the actual opcode bytes like hex
	Flags       OpcodeFlag   // the flags that apply
	Operands    []OpcodeType // Operands like {OP_R64, OP_RM64}
	Description string       // describe
}

func TestFlag(t OpcodeFlag, flag OpcodeFlag) bool {
	return t&flag != 0
}

type OpcodeFlag uint16

const (
	FlagNone  OpcodeFlag = 0
	FlagRexW  OpcodeFlag = 1 << iota // 64-bit operand
	FlagModRM                        // instruction requires ModR/M byte
	FlagRex
	FlagDirect
)

type mnemonic string

var opcodeTable = map[string]OpcodeMeta{
	"MOV_r64_imm64": {
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R64, OP_Imm64},
		Flags:       FlagRexW | FlagRex | FlagDirect,
		Description: "Move a 64 bit immediate into a 64 bit register. (REX.W + 8B /r)",
	},
}

func encode_operation(m mnemonic) {
	util.PackUintLE(0xFF, 8)
}

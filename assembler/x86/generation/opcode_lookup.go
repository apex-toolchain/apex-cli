package generation

import "chasm/util"

type BinaryOrder int

const (
	BaseOpcode BinaryOrder = iota
	AddLow3                // add to the base opcode, like a reg num, so like (opcode + (x & 7))
	UImm                   // unsigned immediate
)

type OpcodeMeta struct {
	Bytes       []byte       // the actual opcode bytes like hex
	Flags       OpcodeFlag   // the flags that apply
	Operands    []OpcodeType // Operands like {OP_R64, OP_RM64}
	Order       []BinaryOrder
	Description string // describe
}

func testFlag(t OpcodeFlag, flag OpcodeFlag) bool {
	return t&flag != 0
}

type OpcodeFlag uint16

const (
	FlagNone  OpcodeFlag = 0
	FlagRexW  OpcodeFlag = 1 << iota // 64-bit operand
	FlagModRM                        // instruction requires ModR/M byte
	FlagRex
	FlagDirect // no mod rm
)

type mnemonic string

var opcodeTable = map[mnemonic]OpcodeMeta{
	"MOV_r64_imm64": {
		Bytes:       []byte{0xB8},
		Operands:    []OpcodeType{OP_R64, OP_Imm64},
		Order:       []BinaryOrder{BaseOpcode, AddLow3, UImm},
		Flags:       FlagRexW | FlagRex | FlagDirect,
		Description: "Move a 64 bit immediate into a 64 bit register. (REX.W + 8B /r)",
	},
}

func EncodeOperation(m mnemonic, input []Operand) []byte {
	meta := opcodeTable[m]
	rex := Rex{
		Needed: false,
		W:      0,
		R:      0,
		X:      0,
		B:      0,
	}

	var out []byte
	if testFlag(meta.Flags, FlagRex) {
		rex.Needed = true
	}

	for _, part := range meta.Order {
		switch part {
		case BaseOpcode:
			out = append(out, meta.Bytes...)
		case AddLow3:
			reg, found := LookupRegCode(input[0].Name)
			if !found {
				panic("reg not found, " + input[0].Name)
			}
			out[len(out)-1] += reg & 7

			if reg&8 > 0 {
				rex.Needed = true
				rex.B = 1
			}

		case UImm:
			out = append(out, util.PackUintLE(input[1].UImm64, 64)...)
		}
	}

	if testFlag(meta.Flags, FlagRexW) {
		rex.Needed = true
		rex.W = 1
	}

	if rex.Needed {
		out = append([]byte{MakeRexByte(rex)}, out...)
	}

	return out
}

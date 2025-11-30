package generation

import "chasm/util"

type BinaryOrder int

const (
	BaseOpcode  BinaryOrder = iota
	AddLow3                 // add to the base opcode, like a reg num, so like (opcode + (x & 7))
	InstantUImm             // unsigned immediate
	Soe0                    //set opcode extension to 0
	SetModNoDisplacement
	DestInRM
	StoreImmediate
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

	"ADD_r64_imm32": {
		Bytes:       []byte{0x81},
		Operands:    []OpcodeType{OP_R64, OP_Imm32},
		Order:       []BinaryOrder{BaseOpcode, SetModNoDisplacement, Soe0, DestInRM, StoreImmediate},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Add a 32 bit immediate to a 64 bit register.",
	},
	"MOV_r64_imm64": {
		Bytes:       []byte{0xB8},
		Operands:    []OpcodeType{OP_R64, OP_Imm64},
		Order:       []BinaryOrder{BaseOpcode, AddLow3, InstantUImm},
		Flags:       FlagRexW | FlagRex | FlagDirect,
		Description: "Move a 64 bit immediate into a 64 bit register. (REX.W + 8B /r)",
	},
}

func EncodeOperation(m mnemonic, input []Operand) []byte {
	meta := opcodeTable[m]

	// Initialize output slice
	var out []byte

	// Initialize REX
	rex := Rex{
		Needed: false,
		W:      0,
		R:      0,
		X:      0,
		B:      0,
	}

	// Initialize ModRM
	modrm := ModRM{
		Needed: testFlag(meta.Flags, FlagModRM),
		Mod:    0,
		RegOp:  0,
		RegMem: 0,
	}

	// Immediate tracking
	var immediate []byte
	var instantImmediate bool // true if it should go immediately after opcode

	// Ensure REX needed if flagged
	if testFlag(meta.Flags, FlagRex) || testFlag(meta.Flags, FlagRexW) {
		rex.Needed = true
		if testFlag(meta.Flags, FlagRexW) {
			rex.W = 1
		}
	}

	// Process the binary order
	for _, part := range meta.Order {
		switch part {
		case BaseOpcode:
			out = append(out, meta.Bytes...)
		case AddLow3:
			regCode, _ := LookupRegCode(input[0].Name)
			out[len(out)-1] += regCode & 7
			if regCode&8 != 0 {
				rex.Needed = true
				rex.B = 1
			}
		case InstantUImm:
			immediate = util.PackUintLE(input[1].UImm64, 64)
			instantImmediate = true
		case StoreImmediate:
			immediate = util.PackUintLE(input[1].UImm64, 32)
			instantImmediate = false
		case SetModNoDisplacement:
			modrm.Mod = MOD_DIRECT_REGISTER
		case Soe0:
			modrm.RegOp = 0
		case DestInRM:
			regCode, _ := LookupRegCode(input[0].Name)
			modrm.RegMem = regCode & 7
			if regCode&8 != 0 {
				rex.Needed = true
				rex.B = 1
			}
		}
	}

	// Append immediate if instant
	if instantImmediate && len(immediate) > 0 {
		out = append(out, immediate...)
	}

	// Append ModRM byte if needed
	if modrm.Needed {
		modrmResult := MakeModRM(modrm.Mod, R_RM{
			DestRegCode: modrm.RegOp,
			SrcRMCode:   modrm.RegMem,
		})
		out = append(out, modrmResult.OutputByte)

		// Update REX if needed from ModRM
		if modrmResult.NeedsREX {
			rex.Needed = true
			rex.R = modrmResult.SetRexR
			rex.B = modrmResult.SetRexB
		}
	}

	// Prepend REX if needed
	if rex.Needed {
		out = append([]byte{MakeRexByte(rex)}, out...)
	}

	// Append normal immediate if it wasn’t instant
	if !instantImmediate && len(immediate) > 0 {
		out = append(out, immediate...)
	}

	return out
}

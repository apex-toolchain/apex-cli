package generation

import "chasm/util"

type BinaryOrder int

const (
	BaseOpcode  BinaryOrder = iota
	AddLow3                 // add to the base opcode, like a reg num, so like (opcode + (x & 7))
	InstantUImm             // unsigned immediate
	Soe0                    //set opcode extension to 0
	SetMod
	SetModNoDisplacement
	DestInRM
	SrcInRM      // set ModRM.RM from second operand (input[1])
	RegOpFromOp0 // set ModRM.reg from first operand (input[0])
	RegOpFromOp1 // set ModRM.reg from second operand (input[1])
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

var opcodeTable = make(map[mnemonic]OpcodeMeta)

// RegisterOpcode registers an opcode entry. This is provided so other
// packages (for example a subpackage that holds instruction files) can
// register metadata without accessing the map directly.
func RegisterOpcode(name string, meta OpcodeMeta) {
	opcodeTable[mnemonic(name)] = meta
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

		case SrcInRM:
			regCode, _ := LookupRegCode(input[1].Name)
			modrm.RegMem = regCode & 7
			if regCode&8 != 0 {
				rex.Needed = true
				rex.B = 1
			}
		case RegOpFromOp0:
			regCode, _ := LookupRegCode(input[0].Name)
			modrm.RegOp = regCode & 7
			if regCode&8 != 0 {
				rex.Needed = true
				rex.R = 1
			}
		case RegOpFromOp1:
			regCode, _ := LookupRegCode(input[1].Name)
			modrm.RegOp = regCode & 7
			if regCode&8 != 0 {
				rex.Needed = true
				rex.R = 1
			}
		}
	}

	if instantImmediate {
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

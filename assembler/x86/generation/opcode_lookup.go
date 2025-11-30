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

var opcodeTable = map[mnemonic]OpcodeMeta{
	"NOP": {
		Bytes:       []byte{0x90},
		Operands:    []OpcodeType{},
		Order:       []BinaryOrder{BaseOpcode},
		Flags:       FlagDirect,
		Description: "Does nothing.",
	},

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
	// MOV r64, r/m64 : dest reg, source rm (0x8B)
	"MOV_r64_rm64": {
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R64, OP_RM64},
		Order:       []BinaryOrder{BaseOpcode, RegOpFromOp0, SrcInRM},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Move r/m64 into r64 (0x8B /r)",
	},
	// MOV r64, r64 : register to register via ModRM (set mod=register)
	"MOV_r64_r64": {
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R64, OP_R64},
		Order:       []BinaryOrder{BaseOpcode, SetModNoDisplacement, RegOpFromOp0, SrcInRM},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Move r64 into r64 using ModRM (register to register) (0x8B /r)",
	},
	// MOV r/m64, r64 : reg to r/m (0x89)
	"MOV_rm64_r64": {
		Bytes:       []byte{0x89},
		Operands:    []OpcodeType{OP_RM64, OP_R64},
		Order:       []BinaryOrder{BaseOpcode, SetModNoDisplacement, RegOpFromOp1, DestInRM},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Move r64 into r/m64 (0x89 /r)",
	},
	// MOV r32, imm32 (B8 + rd)
	"MOV_r32_imm32": {
		Bytes:       []byte{0xB8},
		Operands:    []OpcodeType{OP_R32, OP_Imm32},
		Order:       []BinaryOrder{BaseOpcode, AddLow3, StoreImmediate},
		Flags:       FlagDirect,
		Description: "Move a 32 bit immediate into a 32 bit register (B8+rd)",
	},
	// MOV r32, r/m32 : dest reg, source rm (0x8B)
	"MOV_r32_rm32": {
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R32, OP_RM32},
		Order:       []BinaryOrder{BaseOpcode, RegOpFromOp0, SrcInRM},
		Flags:       FlagModRM,
		Description: "Move r/m32 into r32 (0x8B /r)",
	},
	// MOV r32, r32 : register to register via ModRM (set mod=register)
	"MOV_r32_r32": {
		Bytes:       []byte{0x8B},
		Operands:    []OpcodeType{OP_R32, OP_R32},
		Order:       []BinaryOrder{BaseOpcode, SetModNoDisplacement, RegOpFromOp0, SrcInRM},
		Flags:       FlagModRM,
		Description: "Move r32 into r32 using ModRM (register to register) (0x8B /r)",
	},
	// MOV r/m32, r32 : reg to r/m (0x89)
	"MOV_rm32_r32": {
		Bytes:       []byte{0x89},
		Operands:    []OpcodeType{OP_RM32, OP_R32},
		Order:       []BinaryOrder{BaseOpcode, SetModNoDisplacement, RegOpFromOp1, DestInRM},
		Flags:       FlagModRM,
		Description: "Move r32 into r/m32 (0x89 /r)",
	},
	// MOV r8, r/m8 : dest reg, source rm (0x8A)
	"MOV_r8_rm8": {
		Bytes:       []byte{0x8A},
		Operands:    []OpcodeType{OP_R8, OP_RM8},
		Order:       []BinaryOrder{BaseOpcode, RegOpFromOp0, SrcInRM},
		Flags:       FlagModRM,
		Description: "Move r/m8 into r8 (0x8A /r)",
	},
	// MOV r/m8, r8 : reg to r/m (0x88)
	"MOV_rm8_r8": {
		Bytes:       []byte{0x88},
		Operands:    []OpcodeType{OP_RM8, OP_R8},
		Order:       []BinaryOrder{BaseOpcode, SetModNoDisplacement, RegOpFromOp1, DestInRM},
		Flags:       FlagModRM,
		Description: "Move r8 into r/m8 (0x88 /r)",
	},
	// MOV r/m32, imm32 : C7 /0 (store imm32 in r/m or register)
	"MOV_rm32_imm32": {
		Bytes:       []byte{0xC7},
		Operands:    []OpcodeType{OP_RM32, OP_Imm32},
		Order:       []BinaryOrder{BaseOpcode, Soe0, DestInRM, StoreImmediate},
		Flags:       FlagModRM,
		Description: "Move a 32 bit immediate into r/m32 (C7 /0)",
	},
	// MOV r/m8, imm8 : C6 /0 (store imm8 in r/m or register)
	"MOV_rm8_imm8": {
		Bytes:       []byte{0xC6},
		Operands:    []OpcodeType{OP_RM8, OP_Imm8},
		Order:       []BinaryOrder{BaseOpcode, Soe0, DestInRM, InstantUImm},
		Flags:       FlagModRM,
		Description: "Move an 8 bit immediate into r/m8 (C6 /0)",
	},
	// MOV r8, imm8 (B0 + rd)
	"MOV_r8_imm8": {
		Bytes:       []byte{0xB0},
		Operands:    []OpcodeType{OP_R8, OP_Imm8},
		Order:       []BinaryOrder{BaseOpcode, AddLow3, InstantUImm},
		Flags:       FlagDirect,
		Description: "Move an 8 bit immediate into an 8 bit register (B0+rd)",
	},
	// MOV r/m64, imm32 : C7 /0 (store imm32 in r/m or register)
	"MOV_rm64_imm32": {
		Bytes:       []byte{0xC7},
		Operands:    []OpcodeType{OP_RM64, OP_Imm32},
		Order:       []BinaryOrder{BaseOpcode, Soe0, DestInRM, StoreImmediate},
		Flags:       FlagRexW | FlagRex | FlagModRM,
		Description: "Move a 32 bit immediate into r/m64 (C7 /0)",
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

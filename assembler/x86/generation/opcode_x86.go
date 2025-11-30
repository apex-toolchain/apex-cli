package generation

type RegName string
type RegCode byte

// Integer registers
var regCodes64 = map[RegName]RegCode{
	// 64-bit
	"rax": 0, "rcx": 1, "rdx": 2, "rbx": 3,
	"rsp": 4, "rbp": 5, "rsi": 6, "rdi": 7,
	"r8": 8, "r9": 9, "r10": 10, "r11": 11,
	"r12": 12, "r13": 13, "r14": 14, "r15": 15,

	// 32-bit
	"eax": 0, "ecx": 1, "edx": 2, "ebx": 3,
	"esp": 4, "ebp": 5, "esi": 6, "edi": 7,
	"r8d": 8, "r9d": 9, "r10d": 10, "r11d": 11,
	"r12d": 12, "r13d": 13, "r14d": 14, "r15d": 15,

	// 16-bit
	"ax": 0, "cx": 1, "dx": 2, "bx": 3,
	"sp": 4, "bp": 5, "si": 6, "di": 7,
	"r8w": 8, "r9w": 9, "r10w": 10, "r11w": 11,
	"r12w": 12, "r13w": 13, "r14w": 14, "r15w": 15,

	// 8-bit
	"al": 0, "cl": 1, "dl": 2, "bl": 3,
	"spl": 4, "bpl": 5, "sil": 6, "dil": 7,
	"r8b": 8, "r9b": 9, "r10b": 10, "r11b": 11,
	"r12b": 12, "r13b": 13, "r14b": 14, "r15b": 15,
}

// Floating-point registers (XMM)
var xmmRegs = map[RegName]RegCode{
	"xmm0": 0, "xmm1": 1, "xmm2": 2, "xmm3": 3,
	"xmm4": 4, "xmm5": 5, "xmm6": 6, "xmm7": 7,
	"xmm8": 8, "xmm9": 9, "xmm10": 10, "xmm11": 11,
	"xmm12": 12, "xmm13": 13, "xmm14": 14, "xmm15": 15,
}

func LookupRegCode(name RegName) (byte, bool) {
	if code, ok := regCodes64[name]; ok {
		return byte(code), true
	}
	if code, ok := xmmRegs[name]; ok {
		return byte(code), true
	}
	return 0, false
}

// Operand types
type OperandType int

const (
	Reg OperandType = iota
	Mem
	Imm
)

// Operand struct
type Operand struct {
	Type   OperandType
	Name   RegName
	Disp   int
	UImm64 uint64
	Imm64  int64
}

// Opcode types for instruction metadata
type OpcodeType int

const (
	// R, R/M
	OP_R64 OpcodeType = iota
	OP_RM64
	OP_R32
	OP_RM32
	OP_R16
	OP_RM16
	OP_R8
	OP_RM8

	// Immediates
	OP_Imm8
	OP_Imm16
	OP_Imm32
	OP_Imm64

	// Memory operands
	OP_M64
	OP_M32
	OP_M16
	OP_M8

	// Segment, control, debug (optional)
	OP_SR
	OP_CR
	OP_DR

	// Relative displacements
	OP_REL8
	OP_REL32

	// SIMD
	OP_XMM
	OP_YMM
)

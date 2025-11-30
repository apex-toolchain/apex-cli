package generation

type ModType int

type ModRMResult struct {
	OutputByte byte
	NeedsREX   bool
	RegLow3    byte
	RMLow3     byte
	SetRexR    byte
	SetRexB    byte
}

const (
	MOD_NO_DISPLACEMENT ModType = 0b00
	MOD_8_DISPLACEMENT  ModType = 0b01
	MOD_32_DISPLACEMENT ModType = 0b10
	MOD_DIRECT_REGISTER ModType = 0b11
)

type R_RM struct {
	DestRegCode byte // now takes raw byte instead of RegName
	SrcRMCode   byte // now takes raw byte instead of Operand
}

type ModRM struct {
	Needed bool
	Mod    ModType
	RegOp  byte
	RegMem byte
}

// MakeModRM now works with raw register codes
func MakeModRM(mt ModType, flow R_RM) ModRMResult {
	destCode := flow.DestRegCode
	srcCode := flow.SrcRMCode

	outByte := (srcCode & 7) | ((destCode & 7) << 3) | (byte(mt) << 6)
	needsRex := (destCode|srcCode)&8 != 0

	return ModRMResult{
		OutputByte: outByte,
		NeedsREX:   needsRex,
		RegLow3:    destCode & 7,
		RMLow3:     srcCode & 7,
		SetRexR:    (destCode & 8) >> 3,
		SetRexB:    (srcCode & 8) >> 3,
	}
}

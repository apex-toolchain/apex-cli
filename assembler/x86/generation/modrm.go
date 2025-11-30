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
	DestReg RegName
	SrcRM   Operand
}

type ModRM struct {
	Needed bool // whether to emit a REX byte at all
	Mod    ModType
	RegOp  byte
	RegMem byte
}

func MakeModRM(mt ModType, flow R_RM) ModRMResult {
	destCode, found := LookupRegCode(flow.DestReg)
	if !found {
		panic("Register not found by name: " + flow.DestReg)
	}

	srcCode, found := LookupRegCode(flow.SrcRM.Name)
	if !found {
		panic("Register not found by name: " + flow.SrcRM.Name)
	}

	outByte := (srcCode & 7) | (((destCode >> 3) & 7) << 3) | (byte(mt) << 6)
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

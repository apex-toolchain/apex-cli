package generation

type OperandRole int

const (
	ROLE_DEST_REG OperandRole = iota
	ROLE_SRC_RM
	ROLE_IMM
)

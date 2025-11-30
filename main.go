package main

import (
	"chasm/assembler/x86/generation"
	"fmt"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("chasm: no command provided")
	// 	return
	// }
	// if os.Args[1] == "ld" || os.Args[1] == "link" {
	// 	panic("linking not yet implemented")
	// }

	obs := generation.EncodeOperation("MOV_r64_imm64", []generation.Operand{
		{
			Type: generation.Reg,
			Name: "rsi",
		},
		{
			Type:   generation.Imm,
			UImm64: 1231,
		},
	})
	fmt.Println(obs)
}

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

	obs := generation.EncodeOperation("ADD_r64_imm32", []generation.Operand{
		{
			Type: generation.Reg,
			Name: "r9",
		},
		{
			Type:   generation.Imm,
			UImm64: 124124,
		},
	})
	fmt.Println(obs)

	hexStr := `b"`
	for _, b := range obs {
		hexStr += fmt.Sprintf("\\x%02X", b)
	}
	hexStr += `"`
	fmt.Println(hexStr)
}

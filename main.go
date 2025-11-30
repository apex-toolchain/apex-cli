package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("chasm: no command provided")
		return
	}
	if os.Args[1] == "ld" || os.Args[1] == "link" {
		panic("linking not yet implemented")
	}
}

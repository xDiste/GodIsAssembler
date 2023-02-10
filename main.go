package main

import (
	"debug/elf"
	"fmt"
	"os"
)

// Open and read the ELF file
func readELF(execPath string) *elf.File {
	exec, err := elf.Open(execPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	return exec
}

func main() {
	// Check if there is path of ELF file
	args := os.Args
    	if len(args) == 1 || len(args) > 2 {
		fmt.Println("[*] Usage: patcher /path/of/ELF")
		return
	}
	
	// Read the program
	execPath := args[1]
	exec := readELF(execPath)

	// create the disassembler
	disass := newDisassembler()

	// Disassemble the ELF file
	disass.disassemble(exec)
}

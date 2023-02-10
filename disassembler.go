package main

import (
    "fmt"
	"debug/elf"

	"github.com/knightsc/gapstone"
)

// disassembler class
type disassembler struct{
	capstone_engine *gapstone.Engine
}

// Constructor of disassembler
func newDisassembler() *disassembler{
	return &disassembler{
		capstone_engine: capstoneInit(),
	}
}

func (disass *disassembler) disassemble(exec *elf.File){
	// Find the code section and his size
	codeSection := exec.Section(".text")
	codeSectionLen := codeSection.Size

	// Extract the data
	codeData, _ := codeSection.Data()

	// Find the entrypoint
	//entrypointAddr := exec.Entry

	// Take all the instructions
	assemblyCode, _ := disass.capstone_engine.Disasm(codeData, 0, codeSectionLen)
	
	// Print each instruction in a fancy way
	for i := 0; i < len(assemblyCode); i++ {
		fmt.Printf("0x%x:\t%s\t\t%s\n", assemblyCode[i].Address, assemblyCode[i].Mnemonic, assemblyCode[i].OpStr)
	}
	fmt.Printf("Total number number of instructions: %d\n", len(assemblyCode))	
}

// Init the capstone engine
func capstoneInit() *gapstone.Engine {
	engine, err := gapstone.New(gapstone.CS_ARCH_X86, gapstone.CS_MODE_64)
	if err != nil {
		return nil
	}
	return &engine
}
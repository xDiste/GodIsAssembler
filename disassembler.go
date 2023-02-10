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
	// Take all sections
	sections := exec.Sections
	nInstructions := 0
	for i := 1; i < len(sections); i++ {
		// Take only this sections
		if sections[i].Name == ".init" || sections[i].Name == ".plt" || sections[i].Name == ".text" || sections[i].Name == ".fini" {
			fmt.Printf("Disassembly of section <%s>:\n\n", sections[i].Name)
			currSection := sections[i]
			currSectionLen := sections[i].Size

			// Extract the data
			currSectionData, _ := currSection.Data()

			// Find the start address
			startAddress := currSection.Addr

			// Take all the instructions
			assemblyCode, _ := disass.capstone_engine.Disasm(currSectionData, startAddress, currSectionLen)
			
			// Print each instruction in a fancy way
			for i := 0; i < len(assemblyCode); i++ {
				fmt.Printf("\t0x%x:\t%s\t\t%s\n", assemblyCode[i].Address, assemblyCode[i].Mnemonic, assemblyCode[i].OpStr)
			}
			// Count instructions
			nInstructions += len(assemblyCode)
			fmt.Printf("\n")
		}
	}
	fmt.Printf("Total number number of instructions: %d\n", nInstructions)	
}

// Init the capstone engine
func capstoneInit() *gapstone.Engine {
	engine, err := gapstone.New(gapstone.CS_ARCH_X86, gapstone.CS_MODE_64)
	if err != nil {
		return nil
	}
	return &engine
}
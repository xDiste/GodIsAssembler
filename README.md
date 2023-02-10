# go-disassembler
Just a tiny and simple disassembler using [gapstone](https://github.com/knightsc/gapstone) with Go.

Example using [hello64](./hello64):
```
$ go run . /path/to/ELF
Disassembly of section <.text>:

  0x4000b0:       mov             eax, 1
  0x4000b5:       mov             edi, 1
  0x4000ba:       movabs          rsi, 0xbadc0ff3badc0d3
  0x4000c4:       movabs          rsi, 0x6000e4
  0x4000ce:       mov             edx, 0xc
  0x4000d3:       syscall
  0x4000d5:       mov             eax, 0x3c
  0x4000da:       mov             edi, 0
  0x4000df:       syscall

Total number number of instructions: 9
```

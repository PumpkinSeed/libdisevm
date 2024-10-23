package libdisevm

import (
	"fmt"
	"strings"
)

type Formatter struct {
	Opcodes []Opcode
}

func NewFormatter(opcodes []Opcode) *Formatter {
	return &Formatter{
		Opcodes: opcodes,
	}
}

func (f *Formatter) SingleLine() string {
	var result string
	for _, opcode := range f.Opcodes {
		if opcode.Mnemonic == "" {
			continue
		}
		result += opcode.Mnemonic
		for _, arg := range opcode.Args {
			result += " " + fmt.Sprintf("0x%s", arg)
		}
		result += " "
	}
	return strings.TrimSpace(result)
}

func (f *Formatter) MultiLine() string {
	var result string
	for _, opcode := range f.Opcodes {
		if opcode.Mnemonic == "" {
			continue
		}
		result += opcode.Mnemonic
		for _, arg := range opcode.Args {
			result += " " + fmt.Sprintf("0x%s", arg)
		}
		result += "\n"
	}
	return strings.TrimSpace(result)
}

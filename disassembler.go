package libdisevm

import (
	"errors"
	"strings"
)

type Opcode struct {
	Original []string
	Mnemonic string
	Args     []string
}

func Disassemble(bytecode string) ([]Opcode, error) {
	code := prepareBytecode(bytecode)

	isValidLength := len(code)%2 == 0
	if !isValidLength {
		return nil, errors.New("invalid length")
	}
	opCodeResult := getOperations(code)
	return opCodeResult, nil
}

func prepareBytecode(bytecode string) string {
	// Remove white spaces
	bytecode = strings.TrimSpace(bytecode)

	// Remove prefix
	if strings.HasPrefix(bytecode, "0x") {
		bytecode = bytecode[2:]
	}

	// Transform to uppercase
	bytecode = strings.ToUpper(bytecode)

	return bytecode
}

func getOperations(bytecode string) []Opcode {
	bytecodeSlice := bytecodeToSlice(bytecode)

	var opcodes []Opcode
	for i := 0; i < len(bytecodeSlice); i++ {
		var opcode Opcode
		o := OpcodesMap[bytecodeSlice[i]]
		opcode.Mnemonic = o.Mnemonic
		opcode.Original = append(opcode.Original, bytecodeSlice[i])
		if o.TakeValues != 0 {
			takeValues := uint8(0)
			for takeValues != o.TakeValues {
				i++

				opcode.Args = append(opcode.Args, bytecodeSlice[i])
				opcode.Original = append(opcode.Original, bytecodeSlice[i])
				takeValues++
			}
		}

		opcodes = append(opcodes, opcode)
	}

	return opcodes
}

func bytecodeToSlice(bytecode string) []string {
	var result []string
	for len(bytecode) > 0 {
		result = append(result, bytecode[:2])
		bytecode = bytecode[2:]
	}

	return result
}

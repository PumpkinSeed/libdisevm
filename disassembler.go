package libdisevm

import (
	"errors"
	"strings"
)

func Disassemble(code []byte) string {
	return ""
}

func removeWhiteSpaces(bytecode string) string {
	return strings.TrimSpace(bytecode)
}

func removePrefix(bytecode string) string {
	if strings.HasPrefix(bytecode, "0x") {
		return bytecode[2:]
	}
	return bytecode
}

func isValidLength(bytecode string) bool {
	return len(bytecode)%2 == 0
}

func codeWithUpperCaseLetters(bytecode string) string {
	return strings.ToUpper(bytecode)
}

func getOperations(bytecode string) string {
	opcodeResult := ""
	for len(bytecode) > 0 {
		sliceByteCode := bytecode[:2]
		bytecode = bytecode[2:]
		for _, o := range Opcodes {
			if sliceByteCode == o.Op {
				opcodeResult += o.Mnemonic + " "
				if o.TakeValues != 0 {
					takeValues := uint8(0)
					for takeValues != o.TakeValues {
						value := bytecode[:2]
						bytecode = bytecode[2:]
						opcodeResult += value + " "
						takeValues++
					}
				}
				break
			}
		}
	}
	return opcodeResult
}

func getOpcodes(bytecode string) (string, error) {
	code := removeWhiteSpaces(bytecode)
	codeWithoutPrefix := removePrefix(code)
	codeUpperCaseLetters := codeWithUpperCaseLetters(codeWithoutPrefix)
	if !isValidLength(codeUpperCaseLetters) {
		return "", errors.New("invalid length")
	}
	opCodeResult := getOperations(codeUpperCaseLetters)
	return strings.TrimSpace(opCodeResult), nil
}

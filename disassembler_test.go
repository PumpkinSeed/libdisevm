package libdisevm

import (
	"fmt"
	"testing"
)

func TestDisassembleSingleLine(t *testing.T) {
	data := "604260005260206000F3"
	result, err := Disassemble(data)
	if err != nil {
		t.Error(err)
	}

	formatter := NewFormatter(result)
	expected := "PUSH1 0x42 PUSH1 0x00 MSTORE PUSH1 0x20 PUSH1 0x00 RETURN"
	formatted := formatter.SingleLine()
	if formatted != expected {
		t.Errorf("expected %s, got %s", expected, formatted)
	}
}

func TestDisassembleMultiLine(t *testing.T) {
	data := "604260005260206000F3"
	result, err := Disassemble(data)
	if err != nil {
		t.Error(err)
	}

	formatter := NewFormatter(result)
	formatted := formatter.MultiLine()
	fmt.Println(formatted)
}

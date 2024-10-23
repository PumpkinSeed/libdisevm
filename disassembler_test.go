package libdisevm

import "testing"

func TestDisassemble(t *testing.T) {
	data := "604260005260206000F3"
	result, err := getOpcodes(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

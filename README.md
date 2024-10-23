# libdisevm

EVM bytecode disassembler

### Installation

```shell
go get github.com/PumpkinSeed/libdisevm
```

### Usage

```go
package main

import (
	"fmt"

	"github.com/PumpkinSeed/libdisevm"
)

func main() {
	bytecode := "..."
	opcodes, err := libdisevm.Disassemble(bytecode)
	if err != nil {
		panic(err)
	}

	formatter := libdisevm.NewFormatter(opcodes)

	fmt.Println(formatter.SingleLine())
	// Output: PUSH1 0x42 PUSH1 0x00 MSTORE PUSH1 0x20 PUSH1 0x00 RETURN

	fmt.Println(formatter.MultiLine())
	// Output:
    // PUSH1 0x42
    // PUSH1 0x00
    // MSTORE
    // PUSH1 0x20
    // PUSH1 0x00
    // RETURN
}
```

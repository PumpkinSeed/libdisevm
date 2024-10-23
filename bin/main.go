package main

import (
	"fmt"
	"os"

	"github.com/PumpkinSeed/libdisevm"
)

func main() {
	result, err := libdisevm.Disassemble(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}

	formatter := libdisevm.NewFormatter(result)
	formatted := formatter.Multiline()
	fmt.Println(formatted)
}

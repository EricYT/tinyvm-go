package tdb

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"

	tinyvm "github.com/EricYT/tinyvm-go/tvm"
)

// a debug tool for tinyvm-go, like gdb to c, c++

func Shell(tvm *tinyvm.Ctx) error {
	var running bool = false
	var breakpoints []BreakPoint

	// initialize start register
	tvm.Mem.SetRegisterI32(0x8, tvm.Prog.Start())

	// here we go
	reader := bufio.NewReader(os.Stdin)

LOOP:
	for {
		fmt.Printf("tdb >: ")
		// read input
		input, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err != nil && err == io.EOF {
			break
		}

		// operate args
		tokens := tokenize(input)
		if len(tokens) == 0 {
			fmt.Printf("WARNING: no input.\n")
			continue
		}
		cmd, _ := inputToCmd(string(tokens[0]))

		// cmd operation
		var hitBreakPoint bool
		switch cmd {
		default:
			fmt.Printf("WARNING: \"%s\" is not a valid command.\n", string(tokens[0]))
		case CMD_QUIT:
			break LOOP
		case CMD_RUN:
			if running {
				fmt.Printf("The program is alrady running.\n")
				continue
			}
			hitBreakPoint = run(tvm, breakpoints)
			running = true
		case CMD_BREAK:
			addr := tinyvm.ParseValue(tokens[1])
			bp := BreakPoint{address: addr}
			breakpoints = append(breakpoints, bp)
		case CMD_STEP:
			tvm.Step(tvm.Mem.GetRegisterI32(0x8))
			(*tvm.Mem.GetRegisterI32(0x8))++
			fmt.Printf("Advancing instruction pointer to %d\n", *tvm.Mem.GetRegisterI32(0x8))
		case CMD_CONTINUE:
			tvm.Step(tvm.Mem.GetRegisterI32(0x8))
			(*tvm.Mem.GetRegisterI32(0x8))++
			hitBreakPoint = run(tvm, breakpoints)
		}

		if *tvm.Mem.GetRegisterI32(0x8) > tvm.Prog.NumInstr()-1 {
			fmt.Printf("End of program readched.\n")
			break LOOP
		}

		if hitBreakPoint {
			fmt.Printf("BreakPoint hit at address: %d\n", *tvm.Mem.GetRegisterI32(0x8))
		}
	}

	return nil
}

func run(tvm *tinyvm.Ctx, bps []BreakPoint) bool {
	var idxInstr *int = tvm.Mem.GetRegisterI32(0x8)
	for {
		if *idxInstr > tvm.Prog.NumInstr()-1 {
			return false
		}
		// check break points
		for _, bp := range bps {
			if bp.address == *idxInstr {
				// breakpoint hit
				return true
			}
		}
		// program run one step
		tvm.Step(idxInstr)
		// next step
		(*idxInstr)++
	}
	return false
}

func tokenize(input []byte) [][]byte {
	return bytes.FieldsFunc(input, func(c rune) bool {
		return unicode.IsSpace(c)
	})
}
